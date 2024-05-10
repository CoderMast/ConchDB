<script setup lang="ts">
import router from '../router/index'
import {reactive} from 'vue'
import {Connection} from '../../wailsjs/go/services/Connection'
import {ElMessage} from "element-plus";

// 数据库连接信息
const database = reactive({
  dbtype: "mysql",
  username: "codermast",
  password: "123456",
  dbname: "image_bed",
})

function connection() {
  Connection(database.dbtype, database.username, database.password, database.dbname).then(result => {
    if (result.substring(0, 12) == "Successfully") {
      // 1. 连接成功，弹出提示信息
      ElMessage({
        message: result,
        type: "success"
      })
      // 2. 将页面定向到 Execute 视图
      router.push({name : "Execute"})
    } else {
      ElMessage({
        message: result,
        type: "error"
      })
    }
  })
}

// 数据库类型列表
const dbTypeOptions = [
  {label: "MySQL", value: 'mysql'},
  {label: "Oracle", value: 'oracle'},
  {label: "Other", value: 'other', disabled: true},
];
</script>

<template>
  <el-row>
    <el-col :span="8">数据库类型：</el-col>
    <el-col :span="16">
      <el-select v-model="database.dbtype">
        <el-option
            v-for="item in dbTypeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
        />
      </el-select>
    </el-col>
  </el-row>

  <el-row>
    <el-col :span="8">用户名：</el-col>
    <el-col :span="16">
      <el-input id="username" v-model="database.username" autocomplete="off" type="text"/>
    </el-col>
  </el-row>

  <el-row>
    <el-col :span="8">密码：</el-col>
    <el-col :span="16">
      <el-input id="password" v-model="database.password" autocomplete="off" type="password"/>
    </el-col>
  </el-row>

  <el-row>
    <el-col :span="8">数据库名：</el-col>
    <el-col :span="16">
      <el-input id="databaseName" v-model="database.dbname" autocomplete="off" type="text"/>
    </el-col>
  </el-row>

  <el-button type="info" @click="connection">测试连接</el-button>

</template>

<style scoped>

</style>