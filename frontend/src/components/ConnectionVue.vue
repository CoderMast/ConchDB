<script setup lang="ts">
import router from '../router/index'
import {reactive, watch} from 'vue'
import {Connection} from '../../wailsjs/go/services/Connection'
import {ElMessage} from "element-plus";

// 数据库连接信息
const database = reactive({
  dbtype: "mysql",
  host: "localhost",
  port: "3306",
  username: "codermast",
  password: "123456",
  dbname: "image_bed",
})

// 是否连接成功
let isSuccess = false;

function connection() {
  Connection(database.dbtype, database.host, database.port, database.username, database.password, database.dbname).then(result => {
    if (result.substring(0, 12) == "Successfully") {
      // 1. 连接成功，弹出提示信息
      ElMessage({
        message: result,
        type: "success"
      })
      // 2. 修改连接状态
      isSuccess = true;
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
  {label: "Oracle", value: 'oracle', disabled: true},
  {label: "DM8", value: 'dm8', disabled: true},
  {label: "Other", value: 'other', disabled: true},
];

// 重置
function reset() {
  console.log("reset 被调用")
  Object.assign(database, {
    dbtype: "mysql",
    host: "localhost",
    port: "3306",
    username: "codermast",
    password: "123456",
    dbname: "image_bed",
  })
}

// 确认提交
function submit() {
  if (!isSuccess) {
    ElMessage({
      type: "error",
      message: "请先通过连接测试，成功后方可使用！"
    })
    return;
  }

  ElMessage({
    type: "success",
    message: "连接成功！"
  })

  // 延迟加载
  setTimeout(function () {
  }, 1000);

  router.push({name: "Execute"})
}

// 监视器
watch(
    () => database,
    (newValue, oldValue) => {
      // database 发生变化后，这里被调用
      isSuccess = false;
    }, {
      immediate: false,
      deep: true,
    }
)
</script>

<template>
  <div class="main">
    <h2 class="title">欢迎使用 Conch DBM 管理工具！</h2>

    <el-form label-width="auto" style="max-width: 600px">
      <el-form-item label="数据库类型：">
        <el-select v-model="database.dbtype">
          <el-option
              v-for="item in dbTypeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
              :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="地址(IP)：">
        <el-input id="host" v-model="database.host" autocomplete="off" type="text"/>
      </el-form-item>

      <el-form-item label="端口：">
        <el-input id="port" v-model="database.port" autocomplete="off" type="text"/>
      </el-form-item>

      <el-form-item label="用户名：">
        <el-input id="username" v-model="database.username" autocomplete="off" type="text"/>
      </el-form-item>

      <el-form-item label="密码：">
        <el-input id="password" v-model="database.password" autocomplete="off" type="password"/>
      </el-form-item>

      <el-form-item label="数据库名：">
        <el-input id="databaseName" v-model="database.dbname" autocomplete="off" type="text"/>
      </el-form-item>
    </el-form>
    <div class="buttons-container">
      <el-button type="info" @click="connection">测试连接</el-button>
      <el-button type="danger" @click="reset">清空</el-button>
      <el-button type="primary" @click="submit">确认</el-button>
    </div>
  </div>
</template>

<style scoped>
.title {
  text-align: center;
}

.main {
  position: absolute;
  top: -30vh;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
  width: 500px;
  height: 300px;
}

.buttons-container {
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中（如果需要的话）*/
  /* 如果需要，可以添加一些额外的样式，比如 padding, margin 等 */
  padding: 10px; /* 例如，添加一些内边距 */
}
</style>