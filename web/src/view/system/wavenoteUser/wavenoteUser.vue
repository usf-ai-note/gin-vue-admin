
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="ID" prop="id">
          <el-input v-model.number="searchInfo.id" clearable placeholder="请输入ID" />
        </el-form-item>

        <el-form-item label="邮箱" prop="email">
          <el-input v-model="searchInfo.email" clearable placeholder="请输入邮箱" />
        </el-form-item>

        <el-form-item label="第三方账号" prop="thirdAccount">
          <el-input v-model="searchInfo.thirdAccount" clearable placeholder="请输入第三方账号" />
        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="渠道UID" prop="channelUid">
            <el-input v-model="searchInfo.channelUid" clearable placeholder="请输入渠道UID" />
          </el-form-item>

          <el-form-item label="注册来源" prop="origin">
            <el-select v-model="searchInfo.origin" clearable placeholder="请选择注册来源">
              <el-option :value="1" label="邮箱注册" />
              <el-option :value="2" label="Google ID" />
              <el-option :value="3" label="Apple ID" />
            </el-select>
          </el-form-item>

          <el-form-item label="迁移状态" prop="syncStatus">
            <el-select v-model="searchInfo.syncStatus" clearable placeholder="请选择迁移状态">
              <el-option :value="1" label="迁移" />
              <el-option :value="2" label="被迁移" />
            </el-select>
          </el-form-item>
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        >
            <el-table-column align="left" label="uid" prop="id" width="120" />

            <el-table-column align="left" label="邮箱" prop="email" width="120" />

            <el-table-column align="left" label="第三方账号" prop="thirdAccount" width="120" />

            <el-table-column align="left" label="渠道uid" prop="channelUid" width="120" />

            <el-table-column align="left" label="用户昵称" prop="nickname" width="120" />

            <el-table-column align="left" label="头像" prop="icon" width="120" />

            <el-table-column align="left" label="用户类型" prop="userType" width="120">
    <template #default="scope">{{ formatUserType(scope.row.userType) }}</template>
</el-table-column>

            <el-table-column align="left" label="vip用户的过期时间" prop="vipExpiredAt" width="120" />

            <el-table-column align="left" label="状态" prop="status" width="120">
    <template #default="scope">{{ formatStatus(scope.row.status) }}</template>
</el-table-column>

            <el-table-column align="left" label="注册来源" prop="origin" width="180">
    <template #default="scope">{{ formatOrigin(scope.row.origin) }}</template>
</el-table-column>
            <el-table-column align="left" label="注册时候的系统语言" prop="lan" width="120" />

            <el-table-column align="left" label="转写语言" prop="transLan" width="120" />

            <el-table-column align="left" label="注册国家/地区编码" prop="country" width="120" />

            <el-table-column align="left" label="用户所属行业" prop="industry" width="120" />

            <el-table-column align="left" label="设备ID" prop="deviceId" width="120" />

            <el-table-column align="left" label="注册ip" prop="ip" width="120" />

            <el-table-column align="left" label="迁移状态" prop="syncStatus" width="160">
    <template #default="scope">{{ formatSyncStatus(scope.row.syncStatus) }}</template>
</el-table-column>
            <el-table-column align="left" label="迁移的用户uid" prop="syncUid" width="120" />

            <el-table-column align="left" label="创建时间" prop="createdAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="更新时间" prop="updatedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="删除时间" prop="deletedAt" width="180">
   <template #default="scope">{{ formatDeletedAt(scope.row.deletedAt) }}</template>
</el-table-column>

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="uid">
    {{ detailForm.id }}
</el-descriptions-item>
                    <el-descriptions-item label="邮箱">
    {{ detailForm.email }}
</el-descriptions-item>
                    <el-descriptions-item label="第三方账号">
    {{ detailForm.thirdAccount }}
</el-descriptions-item>
                    <el-descriptions-item label="渠道uid">
    {{ detailForm.channelUid }}
</el-descriptions-item>
                    <el-descriptions-item label="用户昵称">
    {{ detailForm.nickname }}
</el-descriptions-item>
                    <el-descriptions-item label="头像">
    {{ detailForm.icon }}
</el-descriptions-item>
                    <el-descriptions-item label="用户类型">
    {{ formatUserType(detailForm.userType) }}
</el-descriptions-item>
                    <el-descriptions-item label="vip用户的过期时间">
    {{ detailForm.vipExpiredAt }}
</el-descriptions-item>
                    <el-descriptions-item label="状态">
    {{ formatStatus(detailForm.status) }}
</el-descriptions-item>
                    <el-descriptions-item label="注册来源">
    {{ formatOrigin(detailForm.origin) }}
</el-descriptions-item>
                    <el-descriptions-item label="注册时候的系统语言">
    {{ detailForm.lan }}
</el-descriptions-item>
                    <el-descriptions-item label="转写语言">
    {{ detailForm.transLan }}
</el-descriptions-item>
                    <el-descriptions-item label="注册国家/地区编码">
    {{ detailForm.country }}
</el-descriptions-item>
                    <el-descriptions-item label="用户所属行业">
    {{ detailForm.industry }}
</el-descriptions-item>
                    <el-descriptions-item label="设备ID">
    {{ detailForm.deviceId }}
</el-descriptions-item>
                    <el-descriptions-item label="注册ip">
    {{ detailForm.ip }}
</el-descriptions-item>
                    <el-descriptions-item label="迁移状态">
    {{ formatSyncStatus(detailForm.syncStatus) }}
</el-descriptions-item>
                    <el-descriptions-item label="迁移的用户uid">
    {{ detailForm.syncUid }}
</el-descriptions-item>
                    <el-descriptions-item label="创建时间">
    {{ detailForm.createdAt }}
</el-descriptions-item>
                    <el-descriptions-item label="更新时间">
    {{ detailForm.updatedAt }}
</el-descriptions-item>
                    <el-descriptions-item label="删除时间">
    {{ formatDeletedAt(detailForm.deletedAt) }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  findWavenoteUser,
  getWavenoteUserList
} from '@/api/system/wavenoteUser'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ref } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'WavenoteUser'
})

const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

const elSearchFormRef = ref()

const formatOrigin = (value) => {
  const originMap = {
    1: '邮箱注册',
    2: 'Google ID',
    3: 'Apple ID',
  }
  return originMap[value] || value || '-'
}

const formatUserType = (value) => {
  const userTypeMap = {
    0: '普通用户',
    1: 'VIP',
  }
  return userTypeMap[value] || value || '-'
}

const formatStatus = (value) => {
  const statusMap = {
    0: '正常',
    1: '锁定',
  }
  return statusMap[value] || value || '-'
}

const formatSyncStatus = (value) => {
  const syncStatusMap = {
    1: '迁移',
    2: '被迁移',
  }
  return syncStatusMap[value] || value || '-'
}

const formatDeletedAt = (value) => {
  if (!value) {
    return '-'
  }

  const timestamp = Number(value)
  if (Number.isNaN(timestamp)) {
    return value
  }

  // 后端这里返回的是 Unix 秒级时间戳，按 UTC+8 展示。
  const date = new Date((timestamp + 8 * 60 * 60) * 1000)
  const year = date.getUTCFullYear()
  const month = String(date.getUTCMonth() + 1).padStart(2, '0')
  const day = String(date.getUTCDate()).padStart(2, '0')
  const hours = String(date.getUTCHours()).padStart(2, '0')
  const minutes = String(date.getUTCMinutes()).padStart(2, '0')
  const seconds = String(date.getUTCSeconds()).padStart(2, '0')

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
}

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const createSearchInfo = () => ({
  id: undefined,
  email: '',
  thirdAccount: '',
  channelUid: '',
  origin: undefined,
  syncStatus: undefined,
})
const searchInfo = ref(createSearchInfo())
// 重置
const onReset = () => {
  searchInfo.value = createSearchInfo()
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.id === '') {
      searchInfo.value.id = undefined
    }
    if (searchInfo.value.origin === '') {
      searchInfo.value.origin = undefined
    }
    if (searchInfo.value.syncStatus === '') {
      searchInfo.value.syncStatus = undefined
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getWavenoteUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findWavenoteUser({ id: row.id })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
