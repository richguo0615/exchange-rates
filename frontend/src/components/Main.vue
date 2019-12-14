<template>
  <div>
    <div class="wrap">
      <ul class="rate-list">
        <li :style="{width: itemWidth}" v-for="(rate, index) in rateList">
          <div class="list-item">
            <div class="flag">
              <img :src="imageList[rate.Currency]" />
            </div>
            <div class="flag-content">
              <span class="w-currency">{{ rate.Currency }}</span><br/>
              <span class="w-rate">{{ rate.Rate }}</span>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from "axios"
  
export default {
  name: 'Main',
  data () {
    return {
      rateList: [],
      imageList: {
        "USD": "/static/images/USD.png",
        "CNY": "/static/images/CNY.png",
        "JPY": "/static/images/JPY.png",
        "SGD": "/static/images/SGD.png",
        "HKD": "/static/images/HKD.png",
        "EUR": "/static/images/EUR.png",
        "GBP": "/static/images/GBP.png",
      },
      itemWidth: '200px'
    }
  },
  methods: {
    getRateList() {
      let _this = this
      let api = "http://localhost:8080/api/exchangeRates"
      axios.get(api).then(function (res) {
        if (res.status === 200) {
          _this.rateList = res.data
        }
      })
    },
    resizeListWidth( divideNumber ) {
      let wrapWidth = document.querySelector(".rate-list").offsetWidth
      let eachWidth = (wrapWidth / divideNumber) - 15
      this.itemWidth = eachWidth + "px"
    },
    rwdListWidth() {
      let windowWidth = window.innerWidth
      let divideNumber

      if (windowWidth > 1400) {
        divideNumber = 4
      } else if (windowWidth > 1200) {
        divideNumber = 3
      } else if (windowWidth > 900) {
        divideNumber = 2
      } else {
        divideNumber = 2
      }

      this.resizeListWidth( divideNumber )
    }
  },
  created() {
    this.getRateList()
  },
  updated() {
    let _this = this
    _this.rwdListWidth()
    window.onresize = function(event) {
      _this.rwdListWidth()
    };
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped src="@/assets/css/base.css"></style>
