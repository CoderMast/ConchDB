<script setup lang="ts">
import { ref } from 'vue'
import { Execute } from '../../wailsjs/go/services/Connection';
import { FlashOutline } from '@vicons/ionicons5'
import { useMessage } from 'naive-ui'

const message = useMessage();

// 需要执行的 SQL 语句
let SqlStatement = ref<string>("");
let columns = ref();
let data = ref();

async function executeSql() {
  console.log(SqlStatement.value)
  // 调用后端接口执行 SQL
  await Execute(SqlStatement.value).then((result) => {
    message.success("执行成功！")
    let retData = JSON.parse(result)
    columns.value = retData.columns;
    data.value = retData.data.data;
  })
}
</script>

<template>
  <div class="sqlExec">
    <n-input class="" v-model:value="SqlStatement" placeholder="请输入你要执行的 SQL 语句..." type="text">
      <template #prefix>
        <n-icon :component="FlashOutline"/>
      </template>
    </n-input>
    <n-button type="success" @click="executeSql">执行SQL</n-button>
  </div>
  <hr/>
  <n-data-table
      :columns="columns"
      :data="data"
      :bordered="false"
  />
</template>


<style scoped>

.sqlExec {
  max-width: 80vw;
  align-items: center;
}
</style>