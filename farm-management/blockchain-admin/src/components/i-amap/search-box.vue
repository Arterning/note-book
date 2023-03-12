<template>
  <div
    @keydown.up="selectTip('up')"
    @keydown.down="selectTip('down')"
    class="el-vue-search-box-container"
  >
    <div class="search-box-wrapper">
      <input
        v-model="keyword"
        :placeholder="placeholder"
        @keyup.enter="search"
        @input="autoComplete"
        type="text"
      />
      <!-- <span class="search-btn" @click="search">
        <svg-icon class="search" icon-class="search"></svg-icon>
      </span> -->
    </div>
    <div class="search-tips">
      <ul>
        <li
          v-for="(tip, index) in tips"
          :key="index"
          :class="{ 'autocomplete-selected': index === selectedTip }"
          @click="changeTip(tip)"
          @mouseover="selectedTip = index"
        >
          <div>{{ tip.name }}</div>
          <div class="district">{{ tip.district }}</div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { lazyAMapApiLoaderInstance } from 'vue-amap'

export default {
  name: 'ElAmapSearchBox',
  props: ['searchOption', 'onSearchResult', 'events', 'default', 'placeholder'],
  data() {
    return {
      keyword: this.default || '',
      tips: [],
      selectedTip: -1,
      loaded: false
    }
  },
  /* eslint-disable */
  computed: {
    _autoComplete() {
      if (!this.loaded) return
      return new AMap.Autocomplete(this.searchOption || {})
    },
    _placeSearch() {
      if (!this.loaded) return
      return new AMap.PlaceSearch(this.searchOption || {})
    }
  },
  /* eslint-enable */
  mounted() {
    const _loadApiPromise = lazyAMapApiLoaderInstance.load()
    _loadApiPromise.then(() => {
      this.loaded = true
      this._onSearchResult = this.onSearchResult
      // register init event
      this.events &&
        this.events.init &&
        this.events.init({
          autoComplete: this._autoComplete,
          placeSearch: this._placeSearch
        })
    })
  },
  methods: {
    autoComplete() {
      if (!this.keyword || !this._autoComplete) return
      this._autoComplete.search(this.keyword, (status, result) => {
        if (status === 'complete') {
          this.tips = result.tips
        }
      })
    },
    search() {
      this.tips = []
      if (!this._placeSearch) return
      this._placeSearch.search(this.keyword, (status, result) => {
        if (result && result.poiList && result.poiList.count) {
          const {
            poiList: { pois }
          } = result
          const LngLats = pois.map(poi => {
            poi.lat = poi.location.lat
            poi.lng = poi.location.lng
            return poi
          })
          this._onSearchResult(LngLats)
        } else if (result.poiList === undefined) {
          throw new Error(result)
        }
      })
    },
    changeTip(tip) {
      this.keyword = tip.district + tip.name
      this.search()
    },
    selectTip(type) {
      if (type === 'up' && this.selectedTip > 0) {
        this.selectedTip -= 1
        this.keyword = this.tips[this.selectedTip].name
      } else if (type === 'down' && this.selectedTip + 1 < this.tips.length) {
        this.selectedTip += 1
        this.keyword = this.tips[this.selectedTip].name
      }
    }
  }
}
</script>
<style lang="scss" scoped>
.el-vue-search-box-container {
  position: relative;
  width: 360px;
  height: 45px;
  background: #fff;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.15);
  border-radius: 2px 3px 3px 2px;
  z-index: 10;

  .search-box-wrapper {
    position: absolute;
    display: flex;
    align-items: center;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    box-sizing: border-box;

    input {
      flex: 1;
      height: 20px;
      line-height: 20px;
      letter-spacing: 0.5px;
      font-size: 14px;
      text-indent: 10px;
      box-sizing: border-box;
      border: none;
      outline: none;
    }

    .search-btn {
      width: 45px;
      height: 100%;
      display: flex;
      align-items: center;
      justify-content: center;
      background: transparent;
      cursor: pointer;

      .search {
        width: 18px;
        height: 18px;
      }
    }
  }

  .search-tips {
    width: 430px;
    position: absolute;
    top: 40px;
    background: #fff;
    overflow: auto;

    ul {
      padding: 0;
      margin: 0;

      li {
        height: 32px;
        line-height: 32px;
        /*box-shadow: 0 1px 1px rgba(0,0,0,.1);*/
        padding: 0 10px;
        cursor: pointer;
        display: flex;
        font-size: 12px;
        color: #333;

        &.autocomplete-selected {
          background: #eee;
        }

        .district {
          color: #999;
          padding-left: 10px;
        }
      }
    }
  }
}
</style>
