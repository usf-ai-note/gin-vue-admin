
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
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
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="id字段" prop="id" width="120" />

            <el-table-column align="left" label="邮箱" prop="email" width="120" />

            <el-table-column align="left" label="第三方账号" prop="thirdAccount" width="120" />

            <el-table-column align="left" label="渠道uid" prop="channelUid" width="120" />

            <el-table-column align="left" label="密码" prop="pwd" width="120" />

            <el-table-column align="left" label="用户昵称" prop="nickname" width="120" />

            <el-table-column align="left" label="头像" prop="icon" width="120" />

            <el-table-column align="left" label="用户类型;0普通用户;1VIP" prop="userType" width="120" />

            <el-table-column align="left" label="vip用户的过期时间" prop="vipExpiredAt" width="120" />

            <el-table-column align="left" label="0-正常，1-锁定" prop="status" width="120" />

            <el-table-column align="left" label="注册来源:1邮箱注册;2 goole id;3 apple id" prop="origin" width="180">
    <template #default="scope">{{ formatOrigin(scope.row.origin) }}</template>
</el-table-column>
            <el-table-column align="left" label="注册时候的系统语言" prop="lan" width="120" />

            <el-table-column align="left" label="转写语言" prop="transLan" width="120" />

            <el-table-column align="left" label="注册国家/地区编码" prop="country" width="120" />

            <el-table-column align="left" label="用户所属行业" prop="industry" width="120" />

            <el-table-column align="left" label="设备ID" prop="deviceId" width="120" />

            <el-table-column align="left" label="注册ip" prop="ip" width="120" />

            <el-table-column align="left" label="用户迁移状态:1迁移;2被迁移" prop="syncStatus" width="160">
    <template #default="scope">{{ formatSyncStatus(scope.row.syncStatus) }}</template>
</el-table-column>
            <el-table-column align="left" label="迁移的用户uid" prop="syncUid" width="120" />

            <el-table-column align="left" label="创建时间" prop="createdAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="更新时间" prop="updatedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="删除写入时间戳" prop="deletedAt" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateWavenoteUserFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="id字段:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
            <el-form-item label="邮箱:" prop="email">
    <el-input v-model="formData.email" :clearable="true" placeholder="请输入邮箱" />
</el-form-item>
            <el-form-item label="第三方账号:" prop="thirdAccount">
    <el-input v-model="formData.thirdAccount" :clearable="true" placeholder="请输入第三方账号" />
</el-form-item>
            <el-form-item label="渠道uid:" prop="channelUid">
    <el-input v-model="formData.channelUid" :clearable="true" placeholder="请输入渠道uid" />
</el-form-item>
            <el-form-item label="密码:" prop="pwd">
    <el-input v-model="formData.pwd" :clearable="true" placeholder="请输入密码" />
</el-form-item>
            <el-form-item label="用户昵称:" prop="nickname">
    <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入用户昵称" />
</el-form-item>
            <el-form-item label="头像:" prop="icon">
    <el-input v-model="formData.icon" :clearable="true" placeholder="请输入头像" />
</el-form-item>
            <el-form-item label="用户类型;0普通用户;1VIP:" prop="userType">
    <el-input v-model.number="formData.userType" :clearable="true" placeholder="请输入用户类型;0普通用户;1VIP" />
</el-form-item>
            <el-form-item label="vip用户的过期时间:" prop="vipExpiredAt">
    <el-input v-model.number="formData.vipExpiredAt" :clearable="true" placeholder="请输入vip用户的过期时间" />
</el-form-item>
            <el-form-item label="0-正常，1-锁定:" prop="status">
    <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入0-正常，1-锁定" />
</el-form-item>
            <el-form-item label="注册来源:1邮箱注册;2 goole id;3 apple id:" prop="origin">
    <el-select v-model="formData.origin" clearable placeholder="请选择注册来源">
      <el-option :value="1" label="1 - 邮箱注册" />
      <el-option :value="2" label="2 - Google ID" />
      <el-option :value="3" label="3 - Apple ID" />
      <el-option :value="4" label="4" />
    </el-select>
</el-form-item>
            <el-form-item label="注册时候的系统语言:" prop="lan">
    <el-input v-model="formData.lan" :clearable="true" placeholder="请输入注册时候的系统语言" />
</el-form-item>
            <el-form-item label="转写语言:" prop="transLan">
    <el-input v-model="formData.transLan" :clearable="true" placeholder="请输入转写语言" />
</el-form-item>
            <el-form-item label="注册国家/地区编码:" prop="country">
    <el-input v-model="formData.country" :clearable="true" placeholder="请输入注册国家/地区编码" />
</el-form-item>
            <el-form-item label="用户所属行业:" prop="industry">
    <el-input v-model.number="formData.industry" :clearable="true" placeholder="请输入用户所属行业" />
</el-form-item>
            <el-form-item label="设备ID:" prop="deviceId">
    <el-input v-model="formData.deviceId" :clearable="true" placeholder="请输入设备ID" />
</el-form-item>
            <el-form-item label="注册ip:" prop="ip">
    <el-input v-model="formData.ip" :clearable="true" placeholder="请输入注册ip" />
</el-form-item>
            <el-form-item label="用户迁移状态:1迁移;2被迁移:" prop="syncStatus">
    <el-select v-model="formData.syncStatus" clearable placeholder="请选择迁移状态">
      <el-option :value="1" label="1 - 迁移" />
      <el-option :value="2" label="2 - 被迁移" />
    </el-select>
</el-form-item>
            <el-form-item label="迁移的用户uid:" prop="syncUid">
    <el-input v-model.number="formData.syncUid" :clearable="true" placeholder="请输入迁移的用户uid" />
</el-form-item>
            <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="更新时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="删除写入时间戳:" prop="deletedAt">
    <el-input v-model.number="formData.deletedAt" :clearable="true" placeholder="请输入删除写入时间戳" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id字段">
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
                    <el-descriptions-item label="密码">
    {{ detailForm.pwd }}
</el-descriptions-item>
                    <el-descriptions-item label="用户昵称">
    {{ detailForm.nickname }}
</el-descriptions-item>
                    <el-descriptions-item label="头像">
    {{ detailForm.icon }}
</el-descriptions-item>
                    <el-descriptions-item label="用户类型;0普通用户;1VIP">
    {{ detailForm.userType }}
</el-descriptions-item>
                    <el-descriptions-item label="vip用户的过期时间">
    {{ detailForm.vipExpiredAt }}
</el-descriptions-item>
                    <el-descriptions-item label="0-正常，1-锁定">
    {{ detailForm.status }}
</el-descriptions-item>
                    <el-descriptions-item label="注册来源:1邮箱注册;2 goole id;3 apple id">
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
                    <el-descriptions-item label="用户迁移状态:1迁移;2被迁移">
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
                    <el-descriptions-item label="删除写入时间戳">
    {{ detailForm.deletedAt }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createWavenoteUser,
  deleteWavenoteUser,
  deleteWavenoteUserByIds,
  updateWavenoteUser,
  findWavenoteUser,
  getWavenoteUserList
} from '@/api/system/wavenoteUser'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'WavenoteUser'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            email: '',
            thirdAccount: '',
            channelUid: '',
            pwd: '',
            nickname: '',
            icon: '',
            userType: undefined,
            vipExpiredAt: undefined,
            status: undefined,
            origin: undefined,
            lan: '',
            transLan: '',
            country: '',
            industry: undefined,
            deviceId: '',
            ip: '',
            syncStatus: undefined,
            syncUid: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: undefined,
        })



// 验证规则
const rule = reactive({
})

const elFormRef = ref()
const elSearchFormRef = ref()

const formatOrigin = (value) => {
  const originMap = {
    1: '邮箱注册',
    2: 'Google ID',
    3: 'Apple ID',
  }
  return originMap[value] || value || '-'
}

const formatSyncStatus = (value) => {
  const syncStatusMap = {
    1: '迁移',
    2: '被迁移',
  }
  return syncStatusMap[value] || value || '-'
}

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.origin === ""){
        searchInfo.value.origin=null
    }
    if (searchInfo.value.syncStatus === ""){
        searchInfo.value.syncStatus=null
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

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteWavenoteUserFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteWavenoteUserByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWavenoteUserFunc = async(row) => {
    const res = await findWavenoteUser({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWavenoteUserFunc = async (row) => {
    const res = await deleteWavenoteUser({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        id: undefined,
        email: '',
        thirdAccount: '',
        channelUid: '',
        pwd: '',
        nickname: '',
        icon: '',
        userType: undefined,
        vipExpiredAt: undefined,
        status: undefined,
        origin: undefined,
        lan: '',
        transLan: '',
        country: '',
        industry: undefined,
        deviceId: '',
        ip: '',
        syncStatus: undefined,
        syncUid: undefined,
        createdAt: new Date(),
        updatedAt: new Date(),
        deletedAt: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createWavenoteUser(formData.value)
                  break
                case 'update':
                  res = await updateWavenoteUser(formData.value)
                  break
                default:
                  res = await createWavenoteUser(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

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
