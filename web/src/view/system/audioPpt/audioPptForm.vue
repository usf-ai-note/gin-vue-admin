
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createAudioPpt,
  updateAudioPpt,
  findAudioPpt
} from '@/api/system/audioPpt'

defineOptions({
    name: 'AudioPptForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAudioPpt({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
