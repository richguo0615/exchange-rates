<template>
  <div>
    <!--<h1>外幣匯率換算表</h1>-->
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
<style scoped>
div.wrap {
  width: 90%;
  max-width: 1200px;
  height: auto;
  min-height: 300px;
  box-shadow: 0px 0px 3px 3px #dadada;
  margin: 35px auto;
  border-radius: 15px;
  padding: 40px 25px;
  display: flex;
  justify-items: center;
  align-items: center;
}
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  white-space: nowrap;
  padding: 0;
}
li {
  display: inline-block;
  padding: 0 10px;
  float: left;
  border: 1px solid #e0e0e0;
  box-shadow: 0px 0px 2px 1px #dadada;
  background-color: #fff;
  border-radius: 5px;
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  margin-bottom: 20px;
  margin-left: 5px;
  margin-right: 5px;
  min-width: 145px;
}
li .list-item {
  display: flex;
  padding: 0px 0;
}
li .flag {
  width: 35%;
  min-height: 150px;
  display: flex;
  align-items: center;
}
li .flag img {
  max-width: 100%;
}
li .flag-content {
  width: 65%;
  padding: 10px 30px;
  text-align: left;
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
}
li .flag-content span {
  display: inline-block;
  padding: 5px 0px;
  font-weight: bold;
}
a {
  color: #42b983;
}
span.w-currency {
  font-size: 1.5em;
}
span.w-rate {
  font-size: 2em;
}

@media only screen and (max-width: 768px) {
  li .flag, li .flag-content {
    width: 50%;
  }
  li .flag-content {
    padding: 20px 15px;
  }
  span.w-currency {
    font-size: 1em;
  }
  span.w-rate {
    font-size: 1.5em;
  }
}

@media only screen and (max-width: 600px) {
  li .flag {
    width: 40%;
  }
  li .flag-content {
    width: 60%;
  }
  span.w-currency {
    font-size: 1em;
  }
  span.w-rate {
    font-size: 1em;
  }
}

</style>
