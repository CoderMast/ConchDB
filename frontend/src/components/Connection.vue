<script setup lang="ts">
import router from '../router/index'
import { reactive, watch } from 'vue'
import { Connection } from '../../wailsjs/go/services/Connection'
import { useMessage } from 'naive-ui'
import { SQLConnectObject } from '../entity/SQLConnectionObject'

const message = useMessage();

// 数据库连接信息
const sqlConnectObject = reactive<SQLConnectObject>({
  driver: "mysql",
  host: "localhost",
  port: "3306",
  username: "codermast",
  password: "123456",
  database: "image_bed",
})

// 是否连接成功
let isSuccess = false;

function connection() {
  console.log(JSON.stringify(sqlConnectObject))
  Connection(JSON.stringify(sqlConnectObject)).then(result => {
    if (result.substring(0, 12) == "Successfully") {
      // 1. 连接成功，弹出提示信息
      message.success(result);

      // 2. 修改连接状态
      isSuccess = true;

    } else {
      // 连接失败，弹出提示框
      message.error(result)
    }
  })
}

// 数据库类型列表
const dbTypeOptions = [
  { label: "MySQL", value: 'mysql' },
  { label: "Oracle", value: 'oracle', disabled: true },
  { label: "DM8", value: 'dm8', disabled: true },
  { label: "Other", value: 'other', disabled: true },
];

// 重置
function reset() {
  console.log("reset 被调用")
  Object.assign(sqlConnectObject, {
    driver: "mysql",
    host: "localhost",
    port: "3306",
    username: "codermast",
    password: "123456",
    database: "image_bed",
  })
}

// 确认提交
function submit() {
  if (!isSuccess) {
    message.error("请先通过连接测试，成功后方可使用！");
    return;
  }

  message.success("连接成功！");

  // 延迟加载
  setTimeout(function () {
  }, 1000);

  router.push({ name: "Execute" })
}

// 监视器
watch(
  () => sqlConnectObject,
  (newValue, oldValue) => {
    // database 发生变化后，这里被调用
    isSuccess = false;
  }, {
  immediate: false,
  deep: true,
})
</script>

<template>
  <div class="main">
    <h2 class="title">欢迎使用 Conch DBM 管理工具！</h2>

    <!-- 数据库连接表单 -->
    <n-form label-placement="left" label-width="auto" style="max-width: 600px;" require-mark-placement="right-hanging">

      <n-form-item label="数据库类型">
        <n-select v-model:value="sqlConnectObject.driver" :options="dbTypeOptions" placeholder="请选择您的数据库驱动!"
          type="text" />
      </n-form-item>

      <n-form-item label="地址(IP)：">
        <n-input id="host" v-model:value="sqlConnectObject.host" placeholder="请输入数据库的 IP 地址!" type="text" />
      </n-form-item>

      <n-form-item label="端口：">
        <n-input id="port" v-model:value="sqlConnectObject.port" placeholder="请输入数据库的端口!" type="text" />
      </n-form-item>

      <n-form-item label="用户名：">
        <n-input id="username" v-model:value="sqlConnectObject.username" placeholder="请输入数据库用户名!" type="text" />
      </n-form-item>
      <div>
        <n-form-item label="密码：">
          <n-input id="password" v-model:value="sqlConnectObject.password" placeholder="请输入数据库密码!" type="password" />
        </n-form-item>
      </div>
      <n-form-item label="数据库名：">
        <n-input id="databaseName" v-model:value="sqlConnectObject.database" placeholder="请输入数据库名!" type="text" />
      </n-form-item>
    </n-form>


    <!-- 按钮列表 -->
    <div class="buttons-container">
      <n-button class="submitButton" type="info" @click="connection">
        测试连接
      </n-button>
      <n-button class="submitButton" type="error" @click="reset">
        清空
      </n-button>
      <n-button class="submitButton" type="primary" @click="submit">
        确认
      </n-button>
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
  /* 水平居中 */
  justify-content: center;
  /* 垂直居中 */
  align-items: center;
  /* 边距 */
  padding: 10px;
}

.submitButton {
  margin: 10px;
}
</style>