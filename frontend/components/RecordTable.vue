<script lang="ts" setup>
import { RecordItem } from "..";

interface Props {
  records: RecordItem[] | undefined;
}

defineProps<Props>();
</script>

<template>
  <div class="overflow-x-auto">
    <table class="table w-full">
      <thead>
        <tr>
          <th>ID</th>
          <th>Amount</th>
          <th>User Balance</th>
          <th>Operation Response</th>
          <th>Created At</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="record in records">
          <th>{{ record.id }}</th>
          <td>{{ record.amount }}</td>
          <td>{{ record.userBalance }}</td>
          <td>{{ record.operationResponse }}</td>
          <td>{{ record.createdAt }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
