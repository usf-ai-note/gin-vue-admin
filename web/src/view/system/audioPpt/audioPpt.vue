
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

            <el-table-column align="left" label="用户id" prop="uid" width="120" />

            <el-table-column align="left" label="任务唯一识别ID" prop="taskId" width="120" />

            <el-table-column align="left" label="用户id" prop="audioId" width="120" />

            <el-table-column align="left" label="总结结果data_id" prop="summaryId" width="120" />

            <el-table-column align="left" label="PPT 风格" prop="style" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.style) }}</template>
</el-table-column>
            <el-table-column align="left" label="内容详略级别" prop="detailLevel" width="120" />

            <el-table-column align="left" label="页数档位" prop="slideNum" width="120" />

            <el-table-column label="分页提纲 JSON；如果 slide_num 变化，则 outline 需要重新生成" prop="outline" width="200">
   <template #default="scope">
      [富文本内容]
   </template>
</el-table-column>
            <el-table-column align="left" label="模型供应商" prop="provider" width="120" />

            <el-table-column align="left" label="生成状态,300:生成中" prop="status" width="120" />

            <el-table-column align="left" label="图片结果,保存对象存储的path,多张图片逗号分隔" prop="imgs" width="120" />

            <el-table-column align="left" label="幻灯片结果,保存对象存储的path,仅一个" prop="ppt" width="120" />

            <el-table-column align="left" label="pdf结果,保存对象存储的path,仅一个" prop="pdf" width="120" />

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
            <el-table-column align="left" label="删除写入时间戳" prop="deletedAt" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateAudioPptFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="用户id:" prop="uid">
    <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入用户id" />
</el-form-item>
            <el-form-item label="任务唯一识别ID:" prop="taskId">
    <el-input v-model="formData.taskId" :clearable="true" placeholder="请输入任务唯一识别ID" />
</el-form-item>
            <el-form-item label="用户id:" prop="audioId">
    <el-input v-model.number="formData.audioId" :clearable="true" placeholder="请输入用户id" />
</el-form-item>
            <el-form-item label="总结结果data_id:" prop="summaryId">
    <el-input v-model="formData.summaryId" :clearable="true" placeholder="请输入总结结果data_id" />
</el-form-item>
            <el-form-item label="PPT 风格:" prop="style">
    <el-switch v-model="formData.style" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="内容详略级别:" prop="detailLevel">
    <el-input v-model="formData.detailLevel" :clearable="true" placeholder="请输入内容详略级别" />
</el-form-item>
            <el-form-item label="页数档位:" prop="slideNum">
    <el-input v-model="formData.slideNum" :clearable="true" placeholder="请输入页数档位" />
</el-form-item>
            <el-form-item label="分页提纲 JSON；如果 slide_num 变化，则 outline 需要重新生成:" prop="outline">
    <RichEdit v-model="formData.outline"/>
</el-form-item>
            <el-form-item label="模型供应商:" prop="provider">
    <el-input v-model="formData.provider" :clearable="true" placeholder="请输入模型供应商" />
</el-form-item>
            <el-form-item label="生成状态,300:生成中:" prop="status">
    <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入生成状态,300:生成中" />
</el-form-item>
            <el-form-item label="图片结果,保存对象存储的path,多张图片逗号分隔:" prop="imgs">
    <el-input v-model="formData.imgs" :clearable="true" placeholder="请输入图片结果,保存对象存储的path,多张图片逗号分隔" />
</el-form-item>
            <el-form-item label="幻灯片结果,保存对象存储的path,仅一个:" prop="ppt">
    <el-input v-model="formData.ppt" :clearable="true" placeholder="请输入幻灯片结果,保存对象存储的path,仅一个" />
</el-form-item>
            <el-form-item label="pdf结果,保存对象存储的path,仅一个:" prop="pdf">
    <el-input v-model="formData.pdf" :clearable="true" placeholder="请输入pdf结果,保存对象存储的path,仅一个" />
</el-form-item>
            <el-form-item label="幻灯片数量:" prop="slideCount">
    <el-input v-model.number="formData.slideCount" :clearable="true" placeholder="请输入幻灯片数量" />
</el-form-item>
            <el-form-item label="analysis tokens数:" prop="analysisTokens">
    <el-input v-model.number="formData.analysisTokens" :clearable="true" placeholder="请输入analysis tokens数" />
</el-form-item>
            <el-form-item label="image tokens数:" prop="imageTokens">
    <el-input v-model.number="formData.imageTokens" :clearable="true" placeholder="请输入image tokens数" />
</el-form-item>
            <el-form-item label="总tokens数:" prop="totalTokens">
    <el-input v-model.number="formData.totalTokens" :clearable="true" placeholder="请输入总tokens数" />
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
                    <el-descriptions-item label="用户id">
    {{ detailForm.uid }}
</el-descriptions-item>
                    <el-descriptions-item label="任务唯一识别ID">
    {{ detailForm.taskId }}
</el-descriptions-item>
                    <el-descriptions-item label="用户id">
    {{ detailForm.audioId }}
</el-descriptions-item>
                    <el-descriptions-item label="总结结果data_id">
    {{ detailForm.summaryId }}
</el-descriptions-item>
                    <el-descriptions-item label="PPT 风格">
    {{ detailForm.style }}
</el-descriptions-item>
                    <el-descriptions-item label="内容详略级别">
    {{ detailForm.detailLevel }}
</el-descriptions-item>
                    <el-descriptions-item label="页数档位">
    {{ detailForm.slideNum }}
</el-descriptions-item>
                    <el-descriptions-item label="分页提纲 JSON；如果 slide_num 变化，则 outline 需要重新生成">
    <RichView v-model="detailForm.outline" />
</el-descriptions-item>
                    <el-descriptions-item label="模型供应商">
    {{ detailForm.provider }}
</el-descriptions-item>
                    <el-descriptions-item label="生成状态,300:生成中">
    {{ detailForm.status }}
</el-descriptions-item>
                    <el-descriptions-item label="图片结果,保存对象存储的path,多张图片逗号分隔">
    {{ detailForm.imgs }}
</el-descriptions-item>
                    <el-descriptions-item label="幻灯片结果,保存对象存储的path,仅一个">
    {{ detailForm.ppt }}
</el-descriptions-item>
                    <el-descriptions-item label="pdf结果,保存对象存储的path,仅一个">
    {{ detailForm.pdf }}
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
                    <el-descriptions-item label="删除写入时间戳">
    {{ detailForm.deletedAt }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createAudioPpt,
  deleteAudioPpt,
  deleteAudioPptByIds,
  updateAudioPpt,
  findAudioPpt,
  getAudioPptList
} from '@/api/system/audioPpt'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'AudioPpt'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            uid: undefined,
            taskId: '',
            audioId: undefined,
            summaryId: '',
            style: false,
            detailLevel: '',
            slideNum: '',
            outline: '',
            provider: '',
            status: undefined,
            imgs: '',
            ppt: '',
            pdf: '',
            slideCount: undefined,
            analysisTokens: undefined,
            imageTokens: undefined,
            totalTokens: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: undefined,
        })



// 验证规则
const rule = reactive({
})

const elFormRef = ref()
const elSearchFormRef = ref()

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
    if (searchInfo.value.style === ""){
        searchInfo.value.style=null
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
            deleteAudioPptFunc(row)
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
      const res = await deleteAudioPptByIds({ ids })
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
const updateAudioPptFunc = async(row) => {
    const res = await findAudioPpt({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAudioPptFunc = async (row) => {
    const res = await deleteAudioPpt({ id: row.id })
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
        uid: undefined,
        taskId: '',
        audioId: undefined,
        summaryId: '',
        style: false,
        detailLevel: '',
        slideNum: '',
        outline: '',
        provider: '',
        status: undefined,
        imgs: '',
        ppt: '',
        pdf: '',
        slideCount: undefined,
        analysisTokens: undefined,
        imageTokens: undefined,
        totalTokens: undefined,
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
                  res = await createAudioPpt(formData.value)
                  break
                case 'update':
                  res = await updateAudioPpt(formData.value)
                  break
                default:
                  res = await createAudioPpt(formData.value)
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
