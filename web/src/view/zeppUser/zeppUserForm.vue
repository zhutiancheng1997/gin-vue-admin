<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="zepp用户id:" prop="userId">
          <el-input v-model="formData.userId" :clearable="true"  placeholder="请输入zepp用户id" />
       </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true"  placeholder="请输入密码" />
       </el-form-item>
        <el-form-item label="目标步数:" prop="target_step_cnt">
          <el-input v-model.number="formData.target_step_cnt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="今日步数:" prop="todayStepCnt">
          <el-input v-model.number="formData.todayStepCnt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="刷步状态:" prop="pushStatus">
          <el-input v-model="formData.pushStatus" :clearable="true"  placeholder="请输入刷步状态" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createZeppUser,
  updateZeppUser,
  findZeppUser
} from '@/api/zeppUser'

defineOptions({
    name: 'ZeppUserForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            userId: '',
            password: '',
            target_step_cnt: 0,
            todayStepCnt: 0,
            pushStatus: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findZeppUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rezeppUser
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createZeppUser(formData.value)
               break
             case 'update':
               res = await updateZeppUser(formData.value)
               break
             default:
               res = await createZeppUser(formData.value)
               break
           }
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
