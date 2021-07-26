package main

import (
	"context"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/JustHumanz/Go-Simp/pkg/config"
	"github.com/JustHumanz/Go-Simp/pkg/engine"
	network "github.com/JustHumanz/Go-Simp/pkg/network"
	pilot "github.com/JustHumanz/Go-Simp/service/pilot/grpc"
	"github.com/JustHumanz/Go-Simp/service/utility/runfunc"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"

	database "github.com/JustHumanz/Go-Simp/pkg/database"
	log "github.com/sirupsen/logrus"
)

var (
	loc          *time.Location
	Bot          *discordgo.Session
	GroupPayload *[]database.Group
	gRCPconn     = pilot.NewPilotServiceClient(network.InitgRPC(config.Pilot))
)

const (
	ModuleState = "SpaceBiliBili"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableColors: true})
	loc, _ = time.LoadLocation("Asia/Shanghai") /*Use CST*/
}

//Start start twitter module
func main() {
	var (
		configfile config.ConfigFile
		err        error
	)

	GetPayload := func() {
		res, err := gRCPconn.ReqData(context.Background(), &pilot.ServiceMessage{
			Message: "Send me nude",
			Service: ModuleState,
		})
		if err != nil {
			if configfile.Discord != "" {
				pilot.ReportDeadService(err.Error(), ModuleState)
			}
			log.Error("Error when request payload: %s", err)
		}
		err = json.Unmarshal(res.ConfigFile, &configfile)
		if err != nil {
			log.Error(err)
		}

		err = json.Unmarshal(res.VtuberPayload, &GroupPayload)
		if err != nil {
			log.Error(err)
		}
	}

	GetPayload()
	configfile.InitConf()

	Bot, err = discordgo.New("Bot " + configfile.Discord)
	if err != nil {
		log.Error(err)
	}

	database.Start(configfile)

	c := cron.New()
	c.Start()

	c.AddFunc(config.CheckPayload, GetPayload)
	c.AddFunc(config.BiliBiliSpace, CheckSpaceVideo)
	log.Info("Enable " + ModuleState)
	go pilot.RunHeartBeat(gRCPconn, ModuleState)
	runfunc.Run(Bot)
}

//CheckSpaceVideo
func CheckSpaceVideo() {
	for _, GroupData := range *GroupPayload {
		wg := new(sync.WaitGroup)
		for i, MemberData := range GroupData.Members {
			if MemberData.BiliBiliID == 0 && MemberData.Active() {
				continue
			}

			wg.Add(1)
			go func(Group database.Group, Member database.Member, wg *sync.WaitGroup) {
				defer wg.Done()
				log.WithFields(log.Fields{
					"Group":      Group.GroupName,
					"Vtuber":     Member.EnName,
					"BiliBiliID": Member.BiliBiliID,
				}).Info("Checking Space BiliBili")

				Group.RemoveNillIconURL()
				Data := &database.LiveStream{
					Member: Member,
					Group:  Group,
				}
				var (
					Videotype string
					PushVideo engine.SpaceVideo
				)

				body, curlerr := network.CoolerCurl("https://api.bilibili.com/x/space/arc/search?mid="+strconv.Itoa(Data.Member.BiliBiliID)+"&ps="+strconv.Itoa(config.GoSimpConf.LimitConf.SpaceBiliBili), nil)
				if curlerr != nil {
					log.Error(curlerr)
					gRCPconn.ReportError(context.Background(), &pilot.ServiceMessage{
						Message: curlerr.Error(),
						Service: ModuleState,
					})
					return
				}

				err := json.Unmarshal(body, &PushVideo)
				if err != nil {
					log.Error(err)
				}

				for _, video := range PushVideo.Data.List.Vlist {
					if Cover, _ := regexp.MatchString("(?m)(cover|song|feat|music|翻唱|mv|歌曲)", strings.ToLower(video.Title)); Cover || video.Typeid == 31 {
						Videotype = "Covering"
					} else {
						Videotype = "Streaming"
					}

					Data.AddVideoID(video.Bvid).SetType(Videotype).
						UpdateTitle(video.Title).
						UpdateThumbnail(video.Pic).UpdateSchdule(time.Unix(int64(video.Created), 0).In(loc)).
						UpdateViewers(strconv.Itoa(video.Play)).UpdateLength(video.Length).SetState(config.SpaceBili)
					new, id := Data.CheckVideo()
					if new {
						log.WithFields(log.Fields{
							"Vtuber": Data.Member.Name,
						}).Info("New video uploaded")

						Data.InputSpaceVideo()
						video.VideoType = Videotype
						engine.SendLiveNotif(Data, Bot)
					} else {
						if !config.GoSimpConf.LowResources {
							Data.UpdateSpaceViews(id)
						}
					}
				}

			}(GroupData, MemberData, wg)
			if i%config.Waiting == 0 {
				wg.Wait()
			}
		}
		wg.Wait()
	}
}
