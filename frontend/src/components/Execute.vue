<script setup lang="ts">
import {reactive, ref} from 'vue'
import {Execute} from "../../wailsjs/go/services/Connection";

import {ElMessage} from "element-plus";



// 执行 SQL
const sql = ref("select * from users;");
const sqlResult = ref();
const sqlResultHead = ref()
const sqlResultBody = ref()


function execute() {
  // 只要执行，就记录日志
  addSqlExecuteLog();
  // 实际请求
  Execute(sql.value).then(result => {
    if (result.substring(0, 5) != "Error") {
      console.log(result)
      // 将 JSON 字符串转换为二维数组
      const jsonResult = JSON.parse(result);
      console.log(jsonResult)
      // 将二维数组赋值给页面变量
      sqlResult.value = jsonResult;
      sqlResultHead.value = jsonResult[0];
      sqlResultBody.value = jsonResult.slice(1);
      ElMessage(
          {
            message: "执行成功！",
            type: "success"
          }
      )
    } else {
      ElMessage({
        type: "error",
        message: result,
      })
      console.log(result)
    }
  })
}

// SQL 语句执行记录
let executeList = ref(new Array<string>());

// 增加 SQL 执行记录
function addSqlExecuteLog(){
    executeList.value.push(sql.value)
}
</script>

<template>

  <el-row>
    <el-col :span="8">执行SQL</el-col>
    <el-col :span="16">
      <el-input type="text" v-model="sql"/>
    </el-col>
  </el-row>
  <el-button type="primary" @click="execute">执行SQL</el-button>
  <hr/>

  <!-- 表单渲染 -->
  <el-table :data="sqlResultBody" style="width: 100%">
    <el-table-column v-for="item in sqlResultHead" :prop="item" :label="item" />
  </el-table>

  <hr/>
  <!-- 日志渲染 -->
  <el-card style="max-width: 50vw">
    <p align="center" v-for="(item,index) in executeList" :key="index" class="text item">{{ item }}</p>
  </el-card>
</template>

<style scoped>

</style>