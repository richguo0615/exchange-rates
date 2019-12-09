<template>
  <div>
    <h1>外幣匯率換算表</h1>
    <div class="wrap">
      <ul class="rate-list">
        <li v-for="(rate, index) in rateList">
          <div class="list-item">
            <div class="flag">
              <img :src="imageList[rate.Currency]" />
            </div>
            <div class="flag-content">
              <span>{{ rate.Currency }}</span><br/>
              <span>{{ rate.Rate }}</span>
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
      }
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
    }
  },
  created() {
    this.getRateList()
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
  margin: 0 auto;
  border-radius: 15px;
  padding: 25px;
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
  margin: 0 10px;
  float: left;
  border: 1px solid black;
  min-width: 265px;
}
li .list-item {
  display: flex;
  padding: 0px 0;
}
li .flag {
  width: 150px;
  max-width: 200px;
  min-height: 150px;
  display: flex;
  align-items: center;
}
li .flag img {
  max-width: 100%;
}
li .flag-content {
  padding: 10px 30px;
  text-align: left;
}
li .flag-content span {
  display: inline-block;
  padding: 5px 0px;
  font-weight: bold;
}
a {
  color: #42b983;
}
</style>
