package main

import (
	"strconv"
	"strings"

	config "github.com/JustHumanz/Go-Simp/pkg/config"
	database "github.com/JustHumanz/Go-Simp/pkg/database"
	engine "github.com/JustHumanz/Go-Simp/pkg/engine"
	"github.com/bwmarrin/discordgo"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

var (
	Register = &Regis{}
)

func Answer(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.UserID == Register.Admin && m.MessageID == Register.MessageID {
		Register.Start()
		if m.Emoji.MessageFormat() == config.Ok {
			if Register.State == "LiveOnly" {
				Register.SetLiveOnly(true)
			} else if Register.State == "NewUpcoming" {
				Register.SetNewUpcoming(true)
			} else if Register.State == "Dynamic" {
				Register.SetDynamic(true)
			} else if Register.State == "LiteMode" {
				Register.SetLite(true)
			} else if Register.State == "IndieNotif" {
				Register.SetIndieNotif(true)
			}
		} else if m.Emoji.MessageFormat() == config.No {
			if Register.State == "LiveOnly" {
				Register.SetLiveOnly(false)
			} else if Register.State == "NewUpcoming" {
				Register.SetNewUpcoming(false)
			} else if Register.State == "Dynamic" {
				Register.SetDynamic(false)
			} else if Register.State == "LiteMode" {
				Register.SetLite(false)
			} else if Register.State == "IndieNotif" {
				Register.SetIndieNotif(false)
				_, err := s.ChannelMessageSend(m.ChannelID, "[Tips] create a dummy role and tag that role use `"+configfile.BotPrefix.General+"tag roles` command")
				if err != nil {
					log.Error(err)
				}
			}
		}

		Message := func(msg string) {
			_, err := s.ChannelMessageSend(m.ChannelID, msg)
			if err != nil {
				log.Error(err)
			}
		}

		LewdLive := func(def int) {
			_, err := s.ChannelMessageSend(m.ChannelID, "[Error] you can't add livestream with lewd in same channel,canceling lewd")
			if err != nil {
				log.Error(err)
			}
			Register.ChannelState.TypeTag = def
			Register.Clear()
		}

		NillType := func() {
			_, err := s.ChannelMessageSend(m.ChannelID, "[Error] you can't disable all type\nadios")
			if err != nil {
				log.Error(err)
			}
			Register.Clear()
		}

		LiveChange := func() {
			if Register.ChannelState.TypeTag == config.LiveType || Register.ChannelState.TypeTag == config.ArtNLiveType {
				Register.Stop()
				Register.LiveOnly(s)
				Register.BreakPoint(1)

				if !Register.ChannelState.LiveOnly {
					Register.Stop()
					Register.NewUpcoming(s)
					Register.BreakPoint(1)
				}

				Register.Stop()
				Register.Dynamic(s)
				Register.BreakPoint(1)

				Register.Stop()
				Register.Lite(s)
				Register.BreakPoint(1)
			}
		}

		if m.Emoji.MessageFormat() == config.Art {
			if Register.ChannelState.TypeTag == config.LiveType {
				Message("[Info] you enable Livestream and Art type on this channel")
				Register.UpdateType(config.ArtNLiveType)
				err := Register.UpdateChannel("type")
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, err.Error())
					log.Error(err)
				}
				Register.Clear()
				return
			} else if Register.ChannelState.TypeTag == config.LewdType {
				Message("[Info] you enable Lewd and Art type on this channel")
				Register.UpdateType(config.LewdNArtType)
			} else if Register.ChannelState.TypeTag == config.ArtType {
				NillType()
				return
			} else if Register.ChannelState.TypeTag == config.ArtNLiveType {
				Message("[Info] you disable Art type on this channel")
				Register.UpdateType(config.LiveType)
				err := Register.UpdateChannel("type")
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, err.Error())
					log.Error(err)
				}
				Register.Clear()
				return
			} else {
				Message("[Info] you enable Art type on this channel")
				Register.UpdateType(config.ArtType)
			}
		} else if m.Emoji.MessageFormat() == config.Live {
			if Register.ChannelState.TypeTag == config.ArtType {
				Message("[Info] you enable Livestream and Art type on this channel")
				Register.UpdateType(config.ArtNLiveType)
			} else if Register.ChannelState.TypeTag == config.LewdType {
				LewdLive(config.LewdType)
				return
			} else if Register.ChannelState.TypeTag == config.LiveType {
				NillType()
				return
			} else if Register.ChannelState.TypeTag == config.ArtNLiveType {
				Message("[Info] you disable Live type on this channel")
				Register.UpdateType(config.ArtType)
				err := Register.UpdateChannel("type")
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, err.Error())
					log.Error(err)
				}
				return
			} else {
				Message("[Info] you enable Livestream type on this channel")
				Register.UpdateType(config.LiveType)
			}
		} else if m.Emoji.MessageFormat() == config.Lewd {
			if Register.ChannelState.TypeTag == config.LiveType {
				LewdLive(config.LiveType)
				return
			} else if Register.ChannelState.TypeTag == config.ArtType {
				Register.UpdateType(config.LewdNArtType)
			} else if Register.ChannelState.TypeTag == config.ArtNLiveType {
				LewdLive(config.ArtNLiveType)
				return
			} else if Register.ChannelState.TypeTag == config.LewdType {
				NillType()
				return
			} else if Register.ChannelState.TypeTag == config.LewdNArtType {
				Message("[Info] you disable Lewd type on this channel")
				Register.UpdateType(config.ArtType)
			} else {
				Register.UpdateType(config.LewdType)
			}
		}

		if m.Emoji.MessageFormat() == config.One {
			Register.Stop()
			Register.ChoiceType(s)
			Register.BreakPoint(2)

			if Register.ChannelState.Group.GroupName == config.Indie {
				Register.Stop()
				Register.IndieNotif(s)
				Register.BreakPoint(1)
			}

			LiveChange()
			if Register.ChannelState.TypeTag == config.LewdType || Register.ChannelState.TypeTag == config.LewdNArtType {
				if !Register.CheckNSFW(s) {
					return
				}
			}

			err := Register.UpdateChannel("all")
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, err.Error())
				log.Error(err)
			}
			_, err = s.ChannelMessageSend(m.ChannelID, "Done")
			if err != nil {
				log.Error(err)
			}
			Register.Clear()
			return

		} else if m.Emoji.MessageFormat() == config.Two {
			Register.AddRegion(s)

		} else if m.Emoji.MessageFormat() == config.Three {
			Register.DelRegion(s)
		} else if m.Emoji.MessageFormat() == config.Four {
			LiveChange()
			err := Register.UpdateChannel("all")
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, err.Error())
				log.Error(err)
			}
			_, err = s.ChannelMessageSend(m.ChannelID, "Done")
			if err != nil {
				log.Error(err)
			}
			return
		}

		if Register.State == "AddReg" {
			Region := engine.UniCodetoCountryCode(m.Emoji.MessageFormat())
			if Region != "" {
				Register.AddNewRegion(Region)
			}
		} else if Register.State == "DelReg" {
			Region := engine.UniCodetoCountryCode(m.Emoji.MessageFormat())
			if Region != "" {
				Register.RemoveRegion(Region)
			}
		}
	}
}

func RegisterFunc(s *discordgo.Session, m *discordgo.MessageCreate) {
	m.Content = strings.ToLower(m.Content)
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	Prefix := configfile.BotPrefix.General
	Out := func() {
		_, err := s.ChannelMessageSend(m.ChannelID, "Adios")
		if err != nil {
			log.Error(err)
		}
		Register.Clear()
	}

	if strings.HasPrefix(m.Content, Prefix) {
		if m.Content == Prefix+Setup {
			Admin, err := MemberHasPermission(m.GuildID, m.Author.ID)
			if err != nil {
				_, err := s.ChannelMessageSend(m.ChannelID, err.Error())
				if err != nil {
					log.Error(err)
				}
			}

			if Admin {
				_, err := s.ChannelMessageSend(m.ChannelID, "Wellcome to setup mode\ntype `exit` to exit this mode")
				if err != nil {
					log.Error(err)
				}

				Register.SetAdmin(m.Author.ID).SetChannel(m.ChannelID)
				_, err = s.ChannelMessageSend(m.ChannelID, "Select ID or Name of Vtuber group/agency you want to enable (Only one)")
				if err != nil {
					log.Error(err)
				}

				for _, v := range Payload.VtuberData {
					table.Append([]string{strconv.Itoa(int(v.ID)), v.GroupName})
				}
				table.SetHeader([]string{"ID", "GroupName"})
				table.Render()
				_, err = s.ChannelMessageSend(m.ChannelID, "`"+tableString.String()+"`")
				if err != nil {
					log.Error(err)
				}
				Register.UpdateState("Group")
			} else {
				_, err := s.ChannelMessageSend(m.ChannelID, "Your roles don't have permission to enable/disable/update,make sure your roles have `Manage Channels` permission")
				if err != nil {
					log.Error(err)
				}
				Register.Clear()
				return
			}

		} else if m.Content == Prefix+Update2 {
			Admin, err := MemberHasPermission(m.GuildID, m.Author.ID)
			if err != nil {
				_, err := s.ChannelMessageSend(m.ChannelID, err.Error())
				if err != nil {
					log.Error(err)
				}
			}
			if Admin {
				_, err := s.ChannelMessageSend(m.ChannelID, "Wellcome to update mode\ntype `exit` to exit this mode")
				if err != nil {
					log.Error(err)
				}

				var (
					Typestr     string
					LiveOnly    = config.No
					NewUpcoming = config.No
					Dynamic     = config.No
					LiteMode    = config.No
					Indie       = ""
					Region      = "All"
				)
				ChannelData := database.ChannelStatus(m.ChannelID)
				if len(ChannelData) > 0 {
					for _, Channel := range ChannelData {
						if Channel.Region != "" {
							Region = Channel.Region
						}

						if Channel.IsFanart() {
							Typestr = "FanArt"
						}

						if Channel.IsLive() {
							Typestr = "Live"
						}

						if Channel.IsFanart() && Channel.IsLive() {
							Typestr = "FanArt & Livestream"
						}

						if Channel.IsLewd() {
							Typestr = "Lewd"
						}

						if Channel.IsLewd() && Channel.IsFanart() {
							Typestr = "FanArt & Lewd"
						}

						if Channel.LiveOnly {
							LiveOnly = config.Ok
						}

						if Channel.NewUpcoming {
							NewUpcoming = config.Ok
						}

						if Channel.Dynamic {
							Dynamic = config.Ok
						}

						if Channel.LiteMode {
							LiteMode = config.Ok
						}

						if Channel.IndieNotif && Channel.Group.GroupName == config.Indie {
							Indie = config.Ok
						} else if Channel.Group.GroupName != config.Indie {
							Indie = "-"
						} else {
							Indie = config.No
						}

						_, err = s.ChannelMessageSendEmbed(m.ChannelID, engine.NewEmbed().
							SetAuthor(m.Author.Username, m.Author.AvatarURL("128")).
							SetThumbnail(config.GoSimpIMG).
							SetDescription("Channel States of "+Channel.Group.GroupName).
							SetTitle("ID "+strconv.Itoa(int(Channel.ID))).
							AddField("Type", Typestr).
							AddField("LiveOnly", LiveOnly).
							AddField("Dynamic", Dynamic).
							AddField("Upcoming", NewUpcoming).
							AddField("Lite", LiteMode).
							AddField("Regions", Region).
							AddField("Independent notif", Indie).
							InlineAllFields().MessageEmbed)
						if err != nil {
							log.Error(err)
						}
					}
				} else {
					_, err := s.ChannelMessageSendEmbed(m.ChannelID, engine.NewEmbed().
						SetTitle("404 Not found,use `"+Prefix+Setup+"` first").
						SetThumbnail(config.GoSimpIMG).
						SetImage(engine.NotFoundIMG()).MessageEmbed)
					if err != nil {
						log.Error(err)
					}
					return
				}

				Register.SetAdmin(m.Author.ID).SetChannel(m.ChannelID)
				Register.ChannelStates = ChannelData

				_, err = s.ChannelMessageSend(m.ChannelID, "Select ID : ")
				if err != nil {
					log.Error(err)
				}

				Register.UpdateState("SelectChannel")
			} else {
				_, err := s.ChannelMessageSend(m.ChannelID, "You don't have permission to enable/disable/update")
				if err != nil {
					log.Error(err)
				}
				Register.Clear()
				return
			}

		}
	} else if m.Author.ID == Register.Admin && m.ChannelID == Register.ChannelState.ChannelID {
		if m.Content == "exit" {
			Out()
			return
		}
		if Register.State == "SelectChannel" {
			tmp, err := strconv.Atoi(m.Content)
			if err != nil {
				_, err := s.ChannelMessageSend(m.ChannelID, "Worng input ID")
				if err != nil {
					log.Error(err)
				}
				Out()
				return
			} else {
				for _, ChannelState := range Register.ChannelStates {
					if int(ChannelState.ID) == tmp {
						Register.ChannelState = ChannelState
					}
				}
				if Register.ChannelState.ID != 0 {
					Register.SetChannel(m.ChannelID)
					_, err := s.ChannelMessageSend(m.ChannelID, "You selectd `"+Register.ChannelState.Group.GroupName+"` with ID `"+strconv.Itoa(int(Register.ChannelState.ID))+"`")
					if err != nil {
						log.Error(err)
					}
					table.SetHeader([]string{"Menu"})
					table.Append([]string{"Update Channel state"})
					table.Append([]string{"Add region in this channel"})
					table.Append([]string{"Delete region in this channel"})

					if Register.ChannelState.TypeTag == 2 || Register.ChannelState.TypeTag == 3 {
						table.Append([]string{"Change Livestream state"})
					}

					table.Render()
					MsgText, err := s.ChannelMessageSend(m.ChannelID, "`"+tableString.String()+"`")
					if err != nil {
						log.Error(err)
					}

					if Register.ChannelState.TypeTag == 2 || Register.ChannelState.TypeTag == 3 {
						err = engine.Reacting(map[string]string{
							"ChannelID": m.ChannelID,
							"State":     "Menu2",
							"MessageID": MsgText.ID,
						}, s)
						if err != nil {
							log.Error(err)
						}
					} else {
						err = engine.Reacting(map[string]string{
							"ChannelID": m.ChannelID,
							"State":     "Menu",
							"MessageID": MsgText.ID,
						}, s)
						if err != nil {
							log.Error(err)
						}
					}

					Register.UpdateMessageID(MsgText.ID)

				} else {
					_, err := s.ChannelMessageSend(m.ChannelID, "Channel ID not found")
					if err != nil {
						log.Error(err)
					}
					Out()
					return
				}
			}
		}
		if Register.State == "Group" {
			VTuberGroup, err := FindGropName(m.Content)
			if err != nil {
				_, err := s.ChannelMessageSend(m.ChannelID, "`"+m.Content+"`,Name of Vtuber Group was not valid")
				if err != nil {
					log.Error(err)
				}
				return
			}
			Register.SetGroup(VTuberGroup)

			if Register.ChannelState.Group.IsNull() {
				_, err := s.ChannelMessageSendEmbed(m.ChannelID, engine.NewEmbed().
					SetDescription("Invalid ID,Group not found").
					SetImage(engine.NotFoundIMG()).MessageEmbed)
				if err != nil {
					log.Error(err)
				}
				Out()
				return
			}

			Register.Stop()
			for Key, Val := range RegList {
				if Key == Register.ChannelState.Group.GroupName {
					if len(Val) > 3 {
						MsgTxt, err := s.ChannelMessageSend(m.ChannelID, "Select `"+Key+"` region")
						if err != nil {
							log.Error(err)
						}
						Register.UpdateState("AddReg")
						for _, v := range strings.Split(Val, ",") {
							err := s.MessageReactionAdd(m.ChannelID, MsgTxt.ID, engine.CountryCodetoUniCode(v))
							if err != nil {
								log.Error(err)
							}
						}
						Register.UpdateMessageID(MsgTxt.ID)
						Register.BreakPoint(5)

						Register.FixRegion("add")
						if Register.ChannelState.ChannelCheck() {
							_, err := s.ChannelMessageSend(m.ChannelID, "Already setup `"+Register.ChannelState.Group.GroupName+"`,for add/del region use `Update` command")
							if err != nil {
								log.Error(err)
							}
							Out()
							return
						}
						Register.Stop()
					} else {
						Register.RegionTMP = strings.Split(Val, ",")
						Register.FixRegion("add")
					}
				}
			}

			Register.ChoiceType(s)
			Register.BreakPoint(3)

			if Register.ChannelState.TypeTag == 3 || Register.ChannelState.TypeTag == 2 {
				Register.Stop()
				Register.LiveOnly(s)
				Register.BreakPoint(1)

				if !Register.ChannelState.LiveOnly {
					Register.Stop()
					Register.NewUpcoming(s)
					Register.BreakPoint(1)
				}

				Register.Stop()
				Register.Dynamic(s)
				Register.BreakPoint(1)

				Register.Stop()
				Register.Lite(s)
				Register.BreakPoint(1)
			} else if Register.ChannelState.TypeTag == 69 || Register.ChannelState.TypeTag == 70 {
				if !Register.CheckNSFW(s) {
					Out()
					return
				}
			}

			if Register.ChannelState.Group.GroupName == config.Indie {
				Register.Stop()
				Register.IndieNotif(s)
				Register.BreakPoint(1)
			}

			if Register.ChannelState.Group.GroupName != "" {
				err = Register.ChannelState.AddChannel()
				if err != nil {
					log.Error(err)
				}

				_, err = s.ChannelMessageSend(m.ChannelID, "Done,you add `"+Register.ChannelState.Group.GroupName+"` in this channel")
				if err != nil {
					log.Error(err)
				}
			}

			Register.Clear()
			return
		}
	}
}

func (Data *Regis) ChoiceType(s *discordgo.Session) *Regis {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	_, err := s.ChannelMessageSend(Data.ChannelState.ChannelID, "Select Channel Type: ")
	if err != nil {
		log.Error(err)
	}
	Fanart := "Disabled"
	Live := "Disabled"
	Lewd := "Disabled"
	if Data.ChannelState.IsFanart() {
		Fanart = "Enabled"
	}

	if Data.ChannelState.IsLive() {
		Live = "Enabled"
	}

	if Data.ChannelState.IsLewd() {
		Lewd = "Enabled"
	}

	table.SetHeader([]string{"Type", "Select", "Status"})
	table.Append([]string{"Fanart", config.Art, Fanart})
	table.Append([]string{"Livestream", config.Live, Live})
	table.Append([]string{"Lewd Fanart", config.Lewd, Lewd})
	table.Render()

	MsgText, err := s.ChannelMessageSend(Data.ChannelState.ChannelID, "`"+tableString.String()+"`")
	if err != nil {
		s.ChannelMessageSend(Data.ChannelState.ChannelID, "Something error, "+err.Error())
		log.Error(err)
		return nil
	}
	err = engine.Reacting(map[string]string{
		"ChannelID": Data.ChannelState.ChannelID,
		"State":     "TypeChannel",
		"MessageID": MsgText.ID,
	}, s)
	if err != nil {
		log.Error(err)
	}

	Register.UpdateMessageID(MsgText.ID)
	return Data
}

func (Data *Regis) LiveOnly(s *discordgo.Session) *Regis {
	MsgTxt, err := s.ChannelMessageSend(Data.ChannelState.ChannelID, "Enable LiveOnly? **Set livestreams in strict mode(ignoring covering or regular video) notification**")
	if err != nil {
		log.Error(err)
	}
	err = engine.Reacting(map[string]string{
		"ChannelID": Data.ChannelState.ChannelID,
		"State":     "SelectType",
		"MessageID": MsgTxt.ID,
	}, s)
	if err != nil {
		log.Error(err)
	}
	Data.UpdateMessageID(MsgTxt.ID)
	Data.UpdateState("LiveOnly")
	return Data
}

func (Data *Regis) NewUpcoming(s *discordgo.Session) *Regis {
	MsgTxt, err := s.ChannelMessageSend(Data.ChannelState.ChannelID, "Enable NewUpcoming? **Bot will send new upcoming livestream**")
	if err != nil {
		log.Error(err)
	}
	err = engine.Reacting(map[string]string{
		"ChannelID": Data.ChannelState.ChannelID,
		"State":     "SelectType",
		"MessageID": MsgTxt.ID,
	}, s)
	if err != nil {
		log.Error(err)
	}
	Data.UpdateMessageID(MsgTxt.ID)
	Data.UpdateState("NewUpcoming")
	return Data
}

func (Data *Regis) Dynamic(s *discordgo.Session) *Regis {
	MsgTxt, err := s.ChannelMessageSend(Data.ChannelState.ChannelID, "Enable Dynamic mode? **Livestream message will disappear after livestream ended**")
	if err != nil {
		log.Error(err)
	}
	err = engine.Reacting(map[string]string{
		"ChannelID": Data.ChannelState.ChannelID,
		"State":     "SelectType",
		"MessageID": MsgTxt.ID,
	}, s)
	if err != nil {
		log.Error(err)
	}
	Data.UpdateMessageID(MsgTxt.ID)
	Data.UpdateState("Dynamic")
	return Data
}

func (Data *Regis) Lite(s *discordgo.Session) *Regis {
	MsgTxt, err := s.ChannelMessageSend(Data.ChannelState.ChannelID, "Enable Lite mode? **Disabling ping user/role function**")
	if err != nil {
		log.Error(err)
	}
	err = engine.Reacting(map[string]string{
		"ChannelID": Data.ChannelState.ChannelID,
		"State":     "SelectType",
		"MessageID": MsgTxt.ID,
	}, s)
	if err != nil {
		log.Error(err)
	}
	Data.UpdateMessageID(MsgTxt.ID)
	Data.UpdateState("LiteMode")
	return Data
}

func (Data *Regis) IndieNotif(s *discordgo.Session) *Regis {
	MsgTxt, err := s.ChannelMessageSend(Data.ChannelState.ChannelID, "Send all independent vtubers notification?")
	if err != nil {
		log.Error(err)
	}
	err = engine.Reacting(map[string]string{
		"ChannelID": Data.ChannelState.ChannelID,
		"State":     "SelectType",
		"MessageID": MsgTxt.ID,
	}, s)
	if err != nil {
		log.Error(err)
	}
	Data.UpdateMessageID(MsgTxt.ID)
	Data.UpdateState("IndieNotif")
	return Data
}

func (Data *Regis) UpdateChannel(s string) error {
	if s == "type" {
		err := Data.ChannelState.UpdateChannel("Type")
		if err != nil {
			return err
		}
	} else {
		err := Data.ChannelState.UpdateChannel("LiveOnly")
		if err != nil {
			return err
		}

		err = Data.ChannelState.UpdateChannel("Dynamic")
		if err != nil {
			return err
		}

		if !Register.ChannelState.LiveOnly {
			err = Data.ChannelState.UpdateChannel("NewUpcoming")
			if err != nil {
				return err
			}
		}

		err = Data.ChannelState.UpdateChannel("LiteMode")
		if err != nil {
			return err
		}

		if Register.ChannelState.Group.GroupName == config.Indie {
			err = Data.ChannelState.UpdateChannel("IndieNotif")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (Data *Regis) AddRegion(s *discordgo.Session) {
	GroupName := Register.ChannelState.Group.GroupName
	ChannelID := Register.ChannelState.ChannelID

	Register.UpdateState("AddReg")
	RegEmoji := []string{}
	ChannelRegion := strings.Split(Register.ChannelState.Region, ",")
	for _, v := range ChannelRegion {
		if v != "" {
			RegEmoji = append(RegEmoji, engine.CountryCodetoUniCode(v))
			Register.RegionTMP = append(Register.RegionTMP, v)
		}
	}
	_, err := s.ChannelMessageSend(ChannelID, "`"+GroupName+"` Regions you already enabled: "+strings.Join(RegEmoji, "  "))
	if err != nil {
		log.Error(err)
	}

	MsgTxt2, err := s.ChannelMessageSend(ChannelID, "`"+GroupName+" `Select the region you want to add: ")
	if err != nil {
		log.Error(err)
	}

	Register.UpdateMessageID(MsgTxt2.ID)
	for Key, Val := range RegList {
		GroupRegList := strings.Split(Val, ",")

		if Key == GroupName {
			if len(ChannelRegion) == len(GroupRegList) {
				_, err := s.ChannelMessageSend(ChannelID, "You already enable all region")
				if err != nil {
					log.Error(err)
				}
				return
			}

			Register.Stop()
			for _, v2 := range GroupRegList {
				skip := false
				for _, v := range ChannelRegion {
					if v == v2 {
						skip = true
						break
					}
				}
				if !skip {
					err := s.MessageReactionAdd(ChannelID, MsgTxt2.ID, engine.CountryCodetoUniCode(v2))
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
	}

	Register.BreakPoint(3)
	Register.FixRegion("add")
	Register.ChannelState.UpdateChannel("Region")

	_, err = s.ChannelMessageSend(ChannelID, "Done,you added "+strings.Join(Register.AddRegionVal, ","))
	if err != nil {
		log.Error(err)
	}
	Register.Clear()
	return
}

func (Data *Regis) DelRegion(s *discordgo.Session) {
	GroupName := Register.ChannelState.Group.GroupName
	ChannelID := Register.ChannelState.ChannelID

	Register.UpdateState("DelReg")
	RegEmoji := []string{}
	for Key, Val := range RegList {
		if Key == GroupName {
			for _, v2 := range strings.Split(Val, ",") {
				for _, v := range strings.Split(Register.ChannelState.Region, ",") {
					if v == v2 {
						RegEmoji = append(RegEmoji, engine.CountryCodetoUniCode(v2))
						Register.RegionTMP = append(Register.RegionTMP, v2)
					}
				}
			}
		}
	}

	_, err := s.ChannelMessageSend(ChannelID, "`"+GroupName+"` Region you already enabled in here "+strings.Join(RegEmoji, "  "))
	if err != nil {
		log.Error(err)
	}

	MsgID := ""
	if len(RegEmoji) == 0 {
		_, err := s.ChannelMessageSend(ChannelID, "`"+GroupName+"` Region 404,add first")
		if err != nil {
			log.Error(err)
		}
		return
	} else {
		MsgTxt2, err := s.ChannelMessageSend(ChannelID, "`"+GroupName+"` Select region you want to delete : ")
		if err != nil {
			log.Error(err)
		}
		MsgID = MsgTxt2.ID
	}

	Register.Stop()
	for _, v := range RegEmoji {
		err := s.MessageReactionAdd(ChannelID, MsgID, v)
		if err != nil {
			log.Error(err)
		}
	}
	Register.UpdateMessageID(MsgID)
	Register.BreakPoint(4)
	Register.FixRegion("del")
	Register.ChannelState.UpdateChannel("Region")

	_, err = s.ChannelMessageSend(ChannelID, "Done,you remove "+strings.Join(Data.DelRegionVal, ","))
	if err != nil {
		log.Error(err)
	}
	Register.Clear()
	return
}

func (Data *Regis) CheckNSFW(s *discordgo.Session) bool {
	ChannelRaw, err := s.Channel(Data.ChannelState.ChannelID)
	if err != nil {
		log.Error(err)
	}

	if !ChannelRaw.NSFW {
		if Register.ChannelState.TypeTag == 69 {
			_, err := s.ChannelMessageSend(Data.ChannelState.ChannelID, "This Channel was not a NSFW channel")
			if err != nil {
				log.Error(err)
			}
			return false
		} else {
			_, err := s.ChannelMessageSend(Data.ChannelState.ChannelID, "This Channel was not a NSFW channel,change channel type to fanart")
			if err != nil {
				log.Error(err)
			}
			Register.ChannelState.TypeTag = 1
		}

	}
	return true
}
