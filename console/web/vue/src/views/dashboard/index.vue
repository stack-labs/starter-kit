<template>
  <div class="dashboard-container">
    <div class="dashboard-text">name: {{ name }}</div>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="页面" width="150">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="链接">
        <template slot-scope="scope">
          <a @click.prevent="handleLink(scope.row)">{{ scope.row.path }}</a>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import pathToRegexp from "path-to-regexp";

export default {
  name: 'Dashboard',
  computed: {
    ...mapGetters([
      'name'
    ])
  },
  data() {
    return {
      list: [
        {
          'name': 'Gateway',
          'path': '/',
        },
        {
          'name': 'Gateway Metrics',
          'path': '/metrics',
        },
        {
          'name': 'Console首页',
          'path': '/console/',
        },
        {
          'name': 'Echo',
          'path': '/console/v1/echo/',
        },
        {
          'name': 'Gin',
          'path': '/console/v1/gin/',
        },
        {
          'name': 'Iris',
          'path': '/console/v1/iris/',
        },
        {
          'name': 'Beego',
          'path': '/console/v1/beego/',
        },
      ],
      listLoading: false
    }
  },
  methods: {
    handleLink(page) {
      window.open(page.path, '_blank');
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
