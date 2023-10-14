<template>
  <a-config-provider
      :theme="{
        token: {
          colorPrimary: '#00c042',
          colorInfo: '#32b92b',
          borderRadius: 8,
          wireframe: true,
          fontSize: 14,
          sizeStep: 4,
          sizeUnit: 4
        },
      }">
    <router-view/>
  </a-config-provider>
</template>
<script setup>
</script>
<style scoped>
</style>
