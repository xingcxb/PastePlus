<template>
  <a-row>
    <a-card style="padding-top: 30px;width: 620px;height: 300px" :body-style="{padding:0,overflow:'hidden'}">
      <a-flex justify="space-between">
        <img
            alt="avatar"
            src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"
            style="display: block;width:273px"
        />
        <a-flex vertical align="flex-end" justify="space-between" :style="{ padding: '32px' }">
          <a-typography>
            <a-typography-title :level="3">
              “antd is an enterprise-class UI design language and Vue UI library.”
            </a-typography-title>
          </a-typography>
          <a-button type="primary" href="https://antdv.com" target="_blank">Get Start</a-button>
        </a-flex>
      </a-flex>
    </a-card>
  </a-row>
</template>
<script setup>

</script>
<style scoped>

</style>