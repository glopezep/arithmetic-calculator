<script lang="ts" setup>
interface Props {
  show?: boolean;
}

defineProps<Props>();
const emit = defineEmits(["close", "save"]);

const basicOperations = [
  "addition",
  "subtraction",
  "multiplication",
  "division",
];
const firstInput = ref<number | undefined>(undefined);
const secondInput = ref<number | undefined>(undefined);
const currentOperationId = ref<string | undefined>(undefined);

const { data } = useFetch("/api/operations");

const currentOperationType = computed(
  () => data?.value?.items.find((i) => i?.id === currentOperationId.value)?.type
);

watch(currentOperationId, (newValue) => {
  if (!basicOperations.includes(newValue!)) {
    secondInput.value = undefined;
  }

  if (newValue === "random_string") {
    firstInput.value = undefined;
  }
});

async function createOperation() {
  const headers = useRequestHeaders(["cookie"]);

  await $fetch("/api/records", {
    method: "POST",
    headers,
    body: {
      id: currentOperationId.value,
      firstValue: firstInput.value,
      secondValue: secondInput.value,
    },
  });

  emit("save");
}
</script>

<template>
  <div :class="{ modal: true, 'modal-open': show }">
    <div class="modal-box">
      <h3 class="font-bold text-lg py-4">New operation</h3>
      <select
        class="select select-bordered w-full mb-4"
        v-model="currentOperationId"
      >
        <!-- <option :value="undefined">Select</option> -->
        <option v-for="o in data?.items" :value="o.id">{{ o.type }}</option>
      </select>
      <div
        class="grid grid-cols-2 gap-4"
        v-if="
          basicOperations.includes(currentOperationType!)
        "
      >
        <input
          type="number"
          placeholder="Type here"
          class="input input-bordered w-full"
          min="0"
          v-model="firstInput"
        />
        <input
          type="number"
          placeholder="Type here"
          class="input input-bordered w-full"
          v-model="secondInput"
        />
      </div>

      <div
        class="grid grid-cols-2 gap-4"
        v-if="['square_root'].includes(currentOperationType!)"
      >
        <input
          type="number"
          placeholder="Type here"
          class="input input-bordered w-full"
          min="0"
          v-model="firstInput"
        />
      </div>

      <div class="modal-action">
        <button class="btn btn-ghost" @click="$emit('close', $event)">
          Cancel
        </button>
        <button class="btn btn-primary" @click="createOperation">Save</button>
      </div>
    </div>
  </div>
</template>
