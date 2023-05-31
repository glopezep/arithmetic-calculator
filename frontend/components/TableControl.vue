<script lang="ts" setup>
interface Props {
  sortBy: string;
  orderBy: string;
}

defineProps<Props>();
defineEmits(["update:orderBy", "update:sortBy"]);
</script>

<template>
  <div class="form-control mb-4">
    <div class="input-group">
      <span
        @click="$emit('update:orderBy', orderBy === 'desc' ? 'asc' : 'desc')"
      >
        <i
          :class="{
            'i-mdi-sort-ascending': orderBy === 'asc',
            'i-mdi-sort-descending': orderBy === 'desc',
          }"
        ></i>
      </span>
      <select
        class="select select-bordered"
        :value="sortBy"
        @input="$emit('update:sortBy', $event.target?.value)"
      >
        <!-- <option disabl>Select</option> -->
        <option value="user_balance">User balance</option>
        <option value="amount">Amount</option>
        <option value="operation_response">Operation response</option>
        <option value="created_at">Created at</option>
      </select>
    </div>
  </div>
</template>
