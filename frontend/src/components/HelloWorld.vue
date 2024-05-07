<script lang="ts" setup>
import {reactive, ref} from 'vue'
import {Greet,ConnectionMySQL} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "Please enter your Database info 👇",
})

// 数据库连接信息
const database = reactive({
  dbtype: "",
  username: "",
  password: "",
  dbname: "",
  resultText: "",
})

function greet() {
  ConnectionMySQL(database.dbtype,database.username,database.password,database.dbname).then(result => {
    database.resultText = result
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <div>数据库类型：<input id="name" v-model="database.dbtype" autocomplete="off" class="input" type="text"/></div>
      <div>用户名：<input id="name" v-model="database.username" autocomplete="off" class="input" type="text"/></div>
      <div>密码：<input id="name" v-model="database.password" autocomplete="off" class="input" type="text"/></div>
      <div>数据库名：<input id="name" v-model="database.dbname" autocomplete="off" class="input" type="text"/></div>
      <button class="btn" @click="greet">Greet</button>
    </div>
    <div id="result" class="result">{{ database.resultText }}</div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
