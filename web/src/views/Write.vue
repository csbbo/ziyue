<template>
<div id="Write">
  <div class="draft-list">
    <div class="create-item">
      <a class="btn-floating btn-large waves-effect waves-light red"><i class="material-icons">add</i></a>
      <span>创建</span>
    </div>
    <ul>
      <li>aa</li>
      <li>aa</li>
      <li>aa</li>
    </ul>
  </div>

  <div class="write-space">
    <div class="write-bar">
      <button>save</button>
      <button>publish</button>
    </div>
    <input/>
    <textarea></textarea>
  </div>
</div>
</template>

<script>
import "@/less/write.less"
export default {
name: "Write"
}
</script>

<style scoped>

</style>