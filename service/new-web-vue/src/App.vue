<script setup>
import { RouterLink, RouterView } from "vue-router"
import IconHome from "@/components/icons/IconHome.vue"
import "./index.css"
</script>

<template>
  <nav class="nav">
    <div class="navbar">
      <router-link to="/" class="navbar-icon"
        ><IconHome class="h-full" />
        <span class="navbar-icon__span"> Go-Simp </span></router-link
      >
      <div class="navbar-group">
        <a
          href="#"
          class="navbar-toggle navbar-toggle-dark"
          @click="darkMode"
          onclick="return false"
        >
          <font-awesome-icon
            icon="moon"
            class="navbar-toggle-dark__icon fa-fw"
            v-if="theme === 'dark'"
          />
          <font-awesome-icon
            icon="sun"
            class="navbar-toggle-dark__icon fa-fw"
            v-else-if="theme === 'light'"
          />
          <font-awesome-icon
            icon="circle-half-stroke"
            class="navbar-toggle-dark__icon fa-fw"
            v-else
          />
        </a>

        <div class="navbar-menu">
          <a href="#" class="navbar-toggle" onclick="return false">
            <font-awesome-icon
              icon="ellipsis-vertical"
              class="navbar-toggle__icon fa-fw"
            />
          </a>
          <ul class="navbar-items">
            <li class="navbar-item">
              <router-link
                to="/vtubers"
                :class="{ active: isActiveVtuber }"
                class="navbar-link"
                >Vtubers</router-link
              >
            </li>
            <li class="navbar-item">
              <router-link
                to="/docs"
                :class="{ active: isActiveDocs }"
                class="navbar-link"
                >Docs</router-link
              >
            </li>
            <li class="navbar-item">
              <router-link to="/support" class="navbar-link"
                >Support</router-link
              >
            </li>
            <li class="navbar-item">
              <a
                href="https://top.gg/bot/721964514018590802"
                target="_blank"
                class="navbar-link"
                >Vote</a
              >
            </li>
            <li class="navbar-item mobile-menu invite">
              <a
                href="https://discord.com/oauth2/authorize?client_id=721964514018590802&permissions=456720&scope=bot%20applications.commands"
                target="_blank"
                class="navbar-link"
                ><font-awesome-icon
                  icon="right-to-bracket"
                  class="fa-fw mr-1"
                />
                Invite Me</a
              >
            </li>
            <li class="navbar-item mobile-menu dashboard">
              <a
                href="https://web-admin.humanz.moe/login"
                target="_blank"
                class="navbar-link"
                ><font-awesome-icon icon="gauge-simple" class="fa-fw mr-1" />
                Dashboard</a
              >
            </li>
          </ul>
        </div>
      </div>

      <div class="navbar-buttons">
        <a
          class="navbar-buttons__button invite group"
          href="https://discord.com/oauth2/authorize?client_id=721964514018590802&permissions=456720&scope=bot%20applications.commands"
          target="_blank"
          ><font-awesome-icon icon="right-to-bracket" class="fa-fw" />
          <span
            class="navbar-buttons__hover group-hover:!opacity-100 group-hover:!scale-100"
            >Invite Me</span
          >
        </a>
        <a
          class="navbar-buttons__button dashboard group"
          href="https://web-admin.humanz.moe/login"
          target="_blank"
          ><font-awesome-icon icon="gauge-simple" class="fa-fw" />
          <span
            class="navbar-buttons__hover group-hover:!opacity-100 group-hover:!scale-100"
            >Dashboard</span
          >
        </a>
        <!-- Add Dark mode button -->
        <a
          href="#"
          class="navbar-buttons__button dark-mode group"
          @click="darkMode"
          onclick="return false"
        >
          <font-awesome-icon
            icon="moon"
            class="dark-mode-btn__svg fa-fw"
            v-if="theme === 'dark'"
          />
          <font-awesome-icon
            icon="sun"
            class="dark-mode-btn__svg fa-fw"
            v-else-if="theme === 'light'"
          />
          <font-awesome-icon
            icon="circle-half-stroke"
            class="dark-mode-btn__svg fa-fw"
            v-else
          />
          <span
            class="navbar-buttons__hover group-hover:!opacity-100 group-hover:!scale-100"
          >
            {{
              theme === "dark"
                ? "Dark Mode"
                : theme === "light"
                ? "Light Mode"
                : "Auto"
            }}
          </span>
        </a>
      </div>
    </div>
  </nav>

  <main class="mt-[4rem] min-h-[85vh]">
    <RouterView />
  </main>
  <footer class="footer relative z-[1]">
    <div class="footer-page">
      <h2 class="footer-page__title">Vtbot</h2>
      <h4 class="footer-page__subtitle">Bot for Vtubers</h4>
    </div>
    <div class="footer-links">
      <h4 class="footer-links__title">Information</h4>
      <ul class="footer-lists">
        <li class="footer-list">
          <router-link to="/docs" class="footer-list__link">Docs</router-link>
        </li>
        <li class="footer-list">
          <router-link to="/support" class="footer-list__link"
            >Support</router-link
          >
        </li>
        <!-- Github link -->
        <li class="footer-list">
          <a
            href="https://github.com/JustHumanz/Go-Simp"
            class="footer-list__link"
            target="_blank"
          >
            <font-awesome-icon
              :icon="['fab', 'github']"
              class="footer-list__icon fa-fw"
            ></font-awesome-icon>
            Github
          </a>
        </li>
        <!-- Discord link -->
        <li class="footer-list">
          <a
            href="https://discord.gg/ydWC5knbJT"
            target="_blank"
            class="footer-list__link"
          >
            <font-awesome-icon
              :icon="['fab', 'discord']"
              class="footer-list__icon fa-fw"
            ></font-awesome-icon>
            Discord Server
          </a>
        </li>
      </ul>
    </div>
    <div class="footer-credit">
      This Website designed and developed by
      <a
        href="https://github.com/yarndinasti"
        target="_blank"
        class="footer-credit__link"
        >Yarn</a
      >
    </div>
  </footer>
</template>

<script>
// Add Font Awesome for Discord and Github icons
import { library } from "@fortawesome/fontawesome-svg-core"
import { faDiscord, faGithub } from "@fortawesome/free-brands-svg-icons"
import {
  faMoon,
  faSun,
  faCircleHalfStroke,
  faGaugeSimple,
  faEllipsisVertical,
  faRightToBracket,
} from "@fortawesome/free-solid-svg-icons"
import { watch } from "@vue/runtime-core"
library.add(
  faDiscord,
  faGithub,
  faMoon,
  faSun,
  faGaugeSimple,
  faCircleHalfStroke,
  faEllipsisVertical,
  faRightToBracket
)

export default {
  data() {
    return {
      activeListMenu: null,
      activeSubMenu: null,
      menuDocs: false,
      isActiveVtuber: false,
      isActiveDocs: false,
      theme: null,
    }
  },
  async created() {
    this.getClickMenu()
    this.unfocusMenu()

    // window.addEventListener("mousedown", (e) => {
    //   const target = e.target.closest(".navbar-filter-item__link.sub-menu")

    //   if (!target) return

    //   if (this.activeSubMenu === target) {
    //     target.parentElement.parentElement.parentElement.firstElementChild.focus()
    //     this.activeSubMenu = null
    //   } else {
    //     this.activeSubMenu = target
    //   }
    // })

    this.theme = localStorage.getItem("theme")

    this.$watch(
      () => this.$route.params,
      async () => {
        this.isActiveVtuber = this.$route.path.includes("/vtuber")
        this.isActiveDocs = this.$route.path.includes("/docs")
      },

      { immediate: true }
    )

    this.$watch(
      () => this.theme,
      async () => {
        switch (this.theme) {
          case "dark":
            document.documentElement.classList.add("dark")
            break
          case "light":
            document.documentElement.classList.remove("dark")
            break
          default:
            this.autoDark()
        }
      },
      { immediate: true }
    )
  },
  methods: {
    getClickMenu() {
      window.onclick = (e) => {
        if (e.target.id === "router-link") {
          e.preventDefault()
          this.$router.push(e.target.getAttribute("href"))
        }

        if (e.target.closest(".navbar-link")) {
          e.target.closest(".navbar-link").blur()
        }

        if (this.$route.path.includes("/docs")) {
          if (e.target.closest(".docs-menu")) {
            const docsMenu = e.target.closest(".docs-menu")

            this.menuDocs = !this.menuDocs
            if (!this.menuDocs && document.activeElement === docsMenu) {
              docsMenu.blur()
            }
          } else if (e.target.closest(".tab-list__link")) {
            document.activeElement.blur()
            this.menuDocs = false
          } else this.menuDocs = false
        }

        if (!this.$route.path.includes("/vtubers")) {
          this.activeListMenu = null
          return
        }

        if (e.target.closest(".navbar-filter__link")) {
          const navbarFilter = e.target.closest(".navbar-filter__link")

          const liNavbarFilter = navbarFilter.parentElement

          if (liNavbarFilter.classList.contains("disabled")) {
            navbarFilter.blur()
            return
          }

          switch (this.activeListMenu) {
            case navbarFilter:
              this.activeListMenu.blur()
              this.activeListMenu = null
              break
            case null:
              this.activeListMenu = navbarFilter
              break
            default:
              this.activeListMenu = navbarFilter
              break
          }
        } else if (e.target.closest(".navbar-filter-item__link")) {
          const navbarFilterItem = e.target.closest(".navbar-filter-item__link")

          if (!navbarFilterItem.classList.contains("sub-menu")) {
            this.activeListMenu = null
            navbarFilterItem.blur()
          }
        } else if (e.target.closest(".navbar-submenu-item__link")) {
          const navbarSubItem = e.target.closest(".navbar-submenu-item__link")

          this.activeListMenu = null
          navbarSubItem.blur()
        } else if (e.target.closest(".nav-search")) {
          if (this.activeListMenu !== null) {
            console.log("closing menu")
            this.activeListMenu = null
          }

          const navbarSearchItem = e.target.closest(".nav-search")

          navbarSearchItem.children[1].focus()
        } else {
          if (this.activeListMenu === null) return
          console.log("closing menu")
          this.activeListMenu = null
        }
      }
    },
    unfocusMenu() {
      // when document unfocus
      document.onblur = (e) => {
        if (
          this.activeListMenu &&
          this.activeListMenu === document.activeElement
        )
          this.activeListMenu.blur()
        else if (
          !document.activeElement.classList.contains("nav-search__input")
        )
          document.activeElement.blur()

        this.menuDocs = false
        this.activeListMenu = null
      }
    },

    darkMode() {
      switch (this.theme) {
        case "dark":
          localStorage.setItem("theme", "light")
          this.theme = "light"
          break
        case "light":
          localStorage.removeItem("theme")
          this.theme = null
          break
        default:
          localStorage.setItem("theme", "dark")
          this.theme = "dark"
          break
      }
      console.log(this.theme)
    },
    autoDark() {
      if (this.theme) return

      const darkSheme = window.matchMedia("(prefers-color-scheme: dark)")

      if (darkSheme.matches) {
        document.documentElement.classList.add("dark")
      } else {
        document.documentElement.classList.remove("dark")
      }

      darkSheme.onchange = async () => {
        if (darkSheme.matches) {
          document.documentElement.classList.add("dark")
        } else {
          document.documentElement.classList.remove("dark")
        }
      }
    },
  },
}
</script>

<style lang="scss" scoped>
.nav {
  font-family: "Nunito", sans-serif;
  @apply h-16 bg-cyan-500 dark:bg-slate-700 shadow-md shadow-cyan-500/50 dark:shadow-slate-700/50 fixed top-0 left-0 w-screen z-[11] flex justify-center select-none;
  .navbar {
    @apply mx-4 h-full flex justify-between md:justify-around items-center w-full md:w-[90%] lg:w-[85%];

    &-group {
      @apply flex;
    }

    &-menu {
      @apply relative;
    }

    &-icon {
      @apply w-auto h-full p-3 cursor-pointer flex items-center;

      &__span {
        @apply ml-2 text-2xl font-semibold text-center uppercase sm:hidden text-white;
      }
    }
    .navbar-toggle {
      @apply block sm:hidden p-3;

      &__icon,
      &-dark__icon {
        @apply text-white;
      }
    }

    &-menu:focus-within .navbar-items {
      @apply scale-100;
    }

    .navbar-items {
      // make nav-menu items using tailwind
      @apply flex items-start sm:items-center flex-col sm:flex-row bg-cyan-500 dark:bg-slate-600 sm:dark:bg-slate-700 rounded-md sm:bg-transparent w-[13rem] sm:w-auto absolute top-3 right-3 sm:top-auto sm:right-auto sm:static shadow-md shadow-black/50 sm:shadow-none overflow-hidden sm:overflow-visible scale-0 sm:scale-100 origin-top-right transition sm:transition-none duration-200 ease-in-out;
    }

    .navbar-item {
      @apply w-full sm:w-auto sm:mx-1;

      .navbar-link {
        @apply w-full inline-block cursor-pointer text-white py-2 sm:py-1 px-2 sm:rounded-md transition sm:hover:shadow-sm sm:hover:shadow-cyan-600/75 sm:dark:hover:shadow-slate-200/75 hover:bg-black/20 sm:hover:bg-transparent sm:hover:dark:bg-transparent font-semibold hover:sm:-translate-y-px;

        &.router-link-active,
        &.active {
          @apply sm:shadow-md bg-cyan-600 dark:bg-slate-800 sm:bg-transparent dark:sm:bg-transparent sm:shadow-cyan-600/75 sm:dark:shadow-slate-200/75 sm:-translate-y-0.5;
        }
      }

      &.mobile-menu {
        @apply block sm:hidden;

        &.dashboard {
          @apply bg-rose-500 dark:bg-slate-800/50;
        }
        &.invite {
          @apply bg-indigo-500 dark:bg-slate-800/20;
        }
      }
    }

    &-buttons {
      @apply hidden sm:block;

      &__hover {
        @apply absolute top-9 w-max left-1/2 -translate-x-1/2 rounded-md px-2 py-1 shadow-md transition inline-block opacity-0 scale-0;
      }

      &__button {
        @apply my-3 py-1 px-2 rounded-md font-semibold text-white cursor-pointer shadow-none transition-shadow hover:shadow-md relative hover:-translate-y-0.5;

        &.dashboard {
          @apply bg-rose-500 dark:bg-slate-900 hover:shadow-rose-500/75 hover:dark:shadow-slate-100/75 mr-2;

          .navbar-buttons__hover {
            @apply bg-rose-500 dark:bg-slate-900 shadow-rose-500/75 dark:shadow-slate-100/75;
          }
        }
        &.invite {
          @apply bg-indigo-500 dark:bg-slate-800 hover:shadow-indigo-500/75 hover:dark:shadow-slate-100/75 mr-2;

          .navbar-buttons__hover {
            @apply bg-indigo-500 dark:bg-slate-800 shadow-indigo-500/75 dark:shadow-slate-100/75;
          }
        }

        &.dark-mode {
          @apply bg-orange-500 dark:bg-slate-800 hover:shadow-sky-500/75 hover:dark:shadow-slate-100/75;

          .navbar-buttons__hover {
            @apply bg-orange-500 dark:bg-slate-800 shadow-sky-500/75 dark:shadow-slate-100/75;
          }
        }
      }
    }
  }
}
.footer {
  @apply pt-3 bg-cyan-500 dark:bg-slate-700 font-semibold w-full justify-around grid;

  grid-template-columns: 1fr 1fr;
  grid-template-areas:
    "page links"
    "credit credit";

  &-page {
    @apply flex justify-center flex-col sm:px-4 pl-5 md:pl-8 lg:pl-12 xl:pl-20;
    grid-area: page;

    &__title {
      @apply text-3xl font-semibold text-cyan-200 dark:text-slate-300;
    }
    &__subtitle {
      @apply text-white;
    }
  }

  &-links {
    @apply pl-5 md:pl-8 lg:pl-12 xl:pl-20;
    grid-area: links;
    &__title {
      @apply text-xl font-semibold text-cyan-200 dark:text-slate-300;
    }
  }

  &-lists {
    @apply grid;

    @media (min-width: 640px) {
      @apply grid-flow-col;
      grid-template-rows: repeat(2, 1fr);
    }
  }

  &-list {
    @apply w-full;

    &__link {
      @apply text-white hover:text-cyan-100 dark:hover:text-slate-300 inline-block w-full;
    }
  }

  .footer-credit {
    grid-area: credit;
    @apply text-center text-cyan-50 py-2 mt-3 bg-cyan-700 dark:bg-slate-800;
    &__link {
      @apply text-white hover:text-cyan-100 dark:hover:text-slate-300;
    }
  }
}
</style>
