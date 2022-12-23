<template>
  <div>
    <el-table :data="tableData" style="width: 100%;margin-bottom: 20px;" row-key="id" default-expand-all :tree-props="{children: 'children'}">
      <el-table-column prop="name" label="名称" sortable width="180"></el-table-column>
      <el-table-column prop="web_router" label="路由地址" sortable width="180"></el-table-column>
      <el-table-column prop="server_router" label="服务器地址"></el-table-column>
    </el-table>
  </div>
</template>

<script>
import { list } from '@/api/permission'

export default {
  data() {
    return {
      tableData: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      }
    }
  },
  mounted() {
    this.fetchList()
  },
  methods: {
    handleNodeClick(data) {
      console.log(data)
    },
    // 获取列表数据
    fetchList() {
      list().then(response => {
        console.log(response)
        this.tableData = response.data
      })
    }
  }
}
</script>
