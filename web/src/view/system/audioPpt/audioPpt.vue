
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="uid" prop="uid">
          <el-input v-model="searchInfo.uid" clearable placeholder="请输入 uid" style="width: 140px" />
        </el-form-item>
        <el-form-item label="audio_id" prop="audioId">
          <el-input v-model="searchInfo.audioId" clearable placeholder="请输入 audio_id" style="width: 160px" />
        </el-form-item>
        <el-form-item label="生成状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择生成状态" style="width: 160px">
            <el-option
              v-for="item in statusOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

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
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        >
            <el-table-column align="left" label="id" prop="id" width="120" />

            <el-table-column align="left" label="uid" prop="uid" width="120" />

            <el-table-column align="left" label="task id" prop="taskId" width="120" />

            <el-table-column align="left" label="audio_id" prop="audioId" width="120" />

            <el-table-column align="left" label="总结结果data_id" prop="summaryId" width="120" />

            <el-table-column align="left" label="内容详略级别" prop="detailLevel" width="120" />

            <el-table-column align="left" label="页数档位" prop="slideNum" width="120" />

            <el-table-column label="outline" prop="outline" width="200">
   <template #default="scope">
      [富文本内容]
   </template>
</el-table-column>
            <el-table-column align="left" label="模型" prop="provider" width="120" />

            <el-table-column align="left" label="生成状态" prop="status" width="120">
              <template #default="scope">
                {{ formatStatus(scope.row.status) }}
              </template>
            </el-table-column>

            <el-table-column align="left" label="图片path" prop="imgs" width="120" />

            <el-table-column align="left" label="ppt path" prop="ppt" width="120" />

            <el-table-column align="left" label="幻灯片数量" prop="slideCount" width="120" />

            <el-table-column align="left" label="analysis tokens数" prop="analysisTokens" width="120" />

            <el-table-column align="left" label="image tokens数" prop="imageTokens" width="120" />

            <el-table-column align="left" label="总tokens数" prop="totalTokens" width="120" />

            <el-table-column align="left" label="创建时间" prop="createdAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="更新时间" prop="updatedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="删除时间" prop="deletedAt" width="180">
              <template #default="scope">
                {{ formatDeletedAt(scope.row.deletedAt) }}
              </template>
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
                    <el-descriptions-item label="id">
    {{ detailForm.id }}
</el-descriptions-item>
                    <el-descriptions-item label="uid">
    {{ detailForm.uid }}
</el-descriptions-item>
                    <el-descriptions-item label="task id">
    {{ detailForm.taskId }}
</el-descriptions-item>
                    <el-descriptions-item label="audio_id">
    {{ detailForm.audioId }}
</el-descriptions-item>
                    <el-descriptions-item label="总结结果data_id">
    {{ detailForm.summaryId }}
</el-descriptions-item>
                    <el-descriptions-item label="内容详略级别">
    {{ detailForm.detailLevel }}
</el-descriptions-item>
                    <el-descriptions-item label="页数档位">
    {{ detailForm.slideNum }}
</el-descriptions-item>
                    <el-descriptions-item label="outline">
    <RichView v-model="detailForm.outline" />
</el-descriptions-item>
                    <el-descriptions-item label="模型">
    {{ detailForm.provider }}
</el-descriptions-item>
                    <el-descriptions-item label="生成状态">
    {{ formatStatus(detailForm.status) }}
</el-descriptions-item>
                    <el-descriptions-item label="图片path">
    {{ detailForm.imgs }}
</el-descriptions-item>
                    <el-descriptions-item label="ppt path">
    {{ detailForm.ppt }}
</el-descriptions-item>
                    <el-descriptions-item label="幻灯片数量">
    {{ detailForm.slideCount }}
</el-descriptions-item>
                    <el-descriptions-item label="analysis tokens数">
    {{ detailForm.analysisTokens }}
</el-descriptions-item>
                    <el-descriptions-item label="image tokens数">
    {{ detailForm.imageTokens }}
</el-descriptions-item>
                    <el-descriptions-item label="总tokens数">
    {{ detailForm.totalTokens }}
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
  findAudioPpt,
  getAudioPptList
} from '@/api/system/audioPpt'
import RichView from '@/components/richtext/rich-view.vue'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ref } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'AudioPpt'
})

const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const statusOptions = [
  { label: '生成中', value: '300' },
  { label: '生成成功', value: '200' },
  { label: '生成失败', value: '400' }
]
const searchInfo = ref({
  uid: '',
  audioId: '',
  status: undefined
})

const formatStatus = (value) => {
  const match = statusOptions.find(item => item.value === value)
  if (match) {
    return match.label
  }
  if (value === null || value === undefined || value === '') {
    return '-'
  }
  return `未知状态`
}

const formatDeletedAt = (value) => {
  if (value === null || value === undefined || value === '') {
    return '-'
  }
  const numericValue = Number(value)
  if (Number.isNaN(numericValue)) {
    return '-'
  }
  const timestamp = numericValue < 1e12 ? numericValue * 1000 : numericValue
  const date = new Date(timestamp)
  if (Number.isNaN(date.getTime())) {
    return '-'
  }

  const utc8Date = new Date(date.getTime() + 8 * 60 * 60 * 1000)
  const year = utc8Date.getUTCFullYear()
  const month = String(utc8Date.getUTCMonth() + 1).padStart(2, '0')
  const day = String(utc8Date.getUTCDate()).padStart(2, '0')
  const hours = String(utc8Date.getUTCHours()).padStart(2, '0')
  const minutes = String(utc8Date.getUTCMinutes()).padStart(2, '0')
  const seconds = String(utc8Date.getUTCSeconds()).padStart(2, '0')

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
}

// 重置
const onReset = () => {
  searchInfo.value = {
    uid: '',
    audioId: '',
    status: undefined
  }
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
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
  const table = await getAudioPptList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  const res = await findAudioPpt({ id: row.id })
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
