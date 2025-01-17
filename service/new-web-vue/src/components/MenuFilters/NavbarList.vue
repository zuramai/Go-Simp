<script setup>
import GroupsMenu from "./GroupsMenu.vue"
import FilterMenu from "./FilterMenu.vue"
import SortMenu from "./SortMenu.vue"
</script>

<template>
  <nav class="list-nav">
    <ul class="navbar-filters">
      <li
        class="navbar-filter group"
        :class="{
          disabled:
            (disabled && error_status != 404) ||
            (error_status && error_status != 404),
        }"
      >
        <GroupsMenu :groups="groups" />
      </li>
      <li
        class="navbar-filter"
        :class="{
          disabled:
            disabled ||
            error_status ||
            (filters &&
              filters.region.length < 2 &&
              platform.length < 2 &&
              !filters.inactive),
        }"
      >
        <FilterMenu :filters="filters" />
      </li>
      <li class="navbar-filter" :class="{ disabled: disabled || error_status }">
        <SortMenu :filters="filters" />
      </li>
    </ul>
    <div class="nav-search">
      <font-awesome-icon
        icon="magnifying-glass"
        class="fa-fw fa-md nav-search__svg"
      />
      <input
        type="text"
        class="nav-search__input"
        @keydown="searchData"
        ref="search_input"
        :placeholder="placeholder || `Search Vtubers...`"
        :disabled="disable_search"
      />
    </div>
  </nav>
</template>

<script>
import { library } from "@fortawesome/fontawesome-svg-core"
import { faMagnifyingGlass } from "@fortawesome/free-solid-svg-icons"

library.add(faMagnifyingGlass)

export default {
  data() {
    return {
      platform: [],
      search_query: null,
      err_status: null,
    }
  },
  props: {
    groups: {
      type: Array,
      default: [],
    },
    filters: {
      type: Object,
      default: null,
    },
    placeholder: {
      type: String,
      default: "Search Vtubers...",
    },
    disable_search: {
      type: Boolean,
      default: false,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    error_status: {
      type: Number,
      default: null,
    },
  },
  async created() {
    this.$watch(
      () => this.filters,
      () => {
        this.platform = []

        if (this.filters) {
          if (this.filters.youtube) this.platform.push("youtube")
          if (this.filters.twitch) this.platform.push("twitch")
          if (this.filters.bilibili) this.platform.push("bilibili")
        }
      },
      { immediate: true }
    )

    // this.$watch(
    //   () => this.search_query,
    //   () => {
    //     window.scrollTo({
    //       top: 0,
    //     })
    //     this.$emit("search", this.search_query)
    //   },
    //   { immediate: true }
    // )

    this.$watch(
      () => this.$route,
      (a, b) => {
        if (
          a.query.reg !== b.query.reg ||
          a.query.plat !== b.query.plat ||
          a.query.inac !== b.query.inac ||
          a.params.id !== b.params.id
        ) {
          this.$refs.search_input.value = ""
          this.$emit("search", null)
        }
      }
    )
  },
  methods: {
    async searchData() {
      await new Promise((resolve) => setTimeout(resolve, 60))
      this.$emit("search", this.$refs.search_input.value)
    },
  },
}
</script>

<style lang="scss">
.list-nav {
  @apply bg-blue-400 dark:bg-slate-500 fixed top-16 py-2 px-5 w-screen flex flex-wrap-reverse items-center sm:justify-around justify-center z-10 select-none;
}

.navbar-filters {
  @apply flex space-x-1 sm:space-x-2 first:mt-2 xs:first:mt-0 items-center;
}

.navbar-pending {
  @apply w-screen h-[100vw] sm:w-[10.5rem] sm:h-[10.5rem] p-3 flex justify-center items-center;
}

.navbar-filter {
  @apply sm:relative;

  &__link {
    @apply text-white flex space-x-1 items-center font-semibold px-2 py-1 rounded-md transition-all duration-200 ease-in-out;
  }

  &__img {
    @apply min-w-[1.5rem] w-6 object-contain inline-block drop-shadow-md;
  }

  &__svg {
    @apply w-6;
  }

  &__span {
    @apply inline-block xs:hidden;

    @media (min-width: 640px) {
      display: inline-block !important;
    }
  }

  &.disabled {
    @apply opacity-50;

    &__link {
      @apply cursor-not-allowed;
    }
  }

  &:not(.disabled) {
    .navbar-filter__link {
      @apply hover:shadow-sm hover:-translate-y-px shadow-blue-600/75 dark:shadow-slate-300/50;
    }

    &:focus-within {
      .navbar-filter__link {
        @apply shadow-md shadow-blue-600/75 dark:shadow-slate-300/50 -translate-y-0.5;
      }

      .navbar-filter-items {
        @apply scale-y-100;
      }
    }
  }

  &-items {
    @apply absolute flex flex-col bg-blue-400 dark:bg-slate-700 sm:shadow-md sm:shadow-blue-600/75 sm:dark:shadow-slate-200/75 sm:rounded-md transition-all overflow-y-auto overflow-x-hidden max-h-[83.8vh] sm:max-h-60 left-0 mt-2 sm:mt-0 sm:left-auto scale-y-0 origin-top;
    @media (min-width: 640px) {
      scrollbar-width: none; /* Firefox */
      -ms-overflow-style: none; /* IE 10+ */
      &::-webkit-scrollbar {
        /* Chromium and Safari */
        display: none;
      }
    }
  }

  // add class exept sort
  &.group {
    .router-link-active {
      @apply bg-blue-600 dark:bg-slate-900;
    }
  }

  &-item {
    &__img {
      @apply min-w-[1.25rem] w-5 object-contain inline-block drop-shadow-md rounded-md;
    }

    &__svg {
      @apply w-6;
    }

    &__link {
      @apply flex space-x-2 items-center font-semibold px-2 py-1 hover:bg-blue-500/60 dark:hover:bg-slate-900/40 text-white w-screen sm:w-44;

      &.active {
        @apply bg-blue-600 dark:bg-slate-900;
      }

      &.sub-menu::after {
        // add arrow right icon
        @apply content-[''] absolute right-3 rotate-90 transition border-y-5 border-l-5 border-solid border-y-transparent border-l-current;
      }
    }
    &:focus-within {
      .sub-menu::after {
        @apply rotate-0;
      }
      .navbar-submenu-items {
        @apply flex flex-col;
      }
    }
  }
}

.navbar-submenu {
  &-items {
    @apply hidden;
  }

  &-item {
    @apply bg-blue-500/20 dark:bg-slate-900/20 flex items-center w-full;

    &__link {
      @apply text-white flex space-x-2 items-center w-full font-semibold px-2 py-1 hover:bg-blue-500/40 dark:hover:bg-slate-900/40 pl-7 sm:pl-4;

      &.active {
        @apply bg-blue-600 dark:bg-slate-900;
      }
    }

    &__img {
      @apply min-w-[1.25rem] w-5 object-contain inline-block rounded-sm;
    }

    &__svg {
      @apply w-6;
    }
  }
}

.nav-search {
  @apply inline-block mx-1 ml-3 flex-auto sm:flex-none relative hover:-translate-y-px;

  &:focus-within {
    transform: translate(0, -2px) !important;
  }

  &__svg {
    @apply absolute mt-2 ml-2 text-blue-500 dark:text-white;
  }

  &__input {
    @apply bg-blue-300 dark:bg-slate-300/20 focus:bg-blue-200 dark:focus:bg-slate-400 disabled:bg-blue-600 dark:disabled:bg-slate-700 py-1 px-2 rounded-lg transition-all hover:shadow-sm hover:shadow-blue-600/75 dark:hover:shadow-slate-100/75 focus:shadow-md focus:shadow-blue-600/75 dark:focus:shadow-slate-100/75  w-full text-gray-600 dark:text-white font-semibold placeholder:italic placeholder:text-blue-500 dark:placeholder:text-gray-300 disabled:placeholder:text-blue-200 placeholder:font-normal pl-8 focus:outline-none;
  }
}
</style>
