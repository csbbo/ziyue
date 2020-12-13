<template>
<div id="DevTest">
  <!-- Dropdown Trigger -->
  <a class='dropdown-trigger btn' href='#' data-target='dropdown1'>Drop Me!</a>

  <!-- Dropdown Structure -->
  <ul id='dropdown1' class='dropdown-content'>
    <li><a href="#!">one</a></li>
    <li><a href="#!">two</a></li>
    <li class="divider" tabindex="-1"></li>
    <li><a href="#!">three</a></li>
    <li><a href="#!"><i class="material-icons">view_module</i>four</a></li>
    <li><a href="#!"><i class="material-icons">cloud</i>five</a></li>
  </ul>
  <a @click="toast" class="btn">Toast!</a>


  <a class="btn-floating btn-large waves-effect waves-light red"><i class="material-icons">add</i></a>
</div>
</template>

<script>
import { serverError } from "@/common/utils";
export default {
name: "DevTest",
  mounted() {
    window.jQuery('.dropdown-trigger').dropdown()
  },
  methods: {
    toast() {
      serverError()
    }
  }
}
</script>

<style scoped>

</style>