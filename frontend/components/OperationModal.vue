<script lang="ts" setup>
interface Props {
  show?: boolean;
}

defineProps<Props>();

const emit = defineEmits<{
  (e: "close", event: MouseEvent): void;
}>();

const basicOperations = [
  "addition",
  "subtraction",
  "multiplication",
  "division",
];
const firstInput = ref<number | undefined>(undefined);
const secondInput = ref<number | undefined>(undefined);
const operationType = ref<string>("addition");

watch(operationType, (newValue) => {
  if (!basicOperations.includes(newValue)) {
    secondInput.value = undefined;
  }

  if (newValue === "random_string") {
    firstInput.value = undefined;
  }
});

function createOperation() {
  const headers = useRequestHeaders(["cookie"]);

  return $fetch("/api/operations", {
    method: "POST",
    headers,
    body: {
      type: operationType.value,
      firstInput: firstInput.value,
      secondInput: secondInput.value,
    },
  });
}
</script>

<template>
  <div :class="{ modal: true, 'modal-open': show }">
    <div class="modal-box">
      <h3 class="font-bold text-lg py-4">New operation</h3>
      <select
        class="select select-bordered w-full mb-4"
        v-model="operationType"
      >
        <option value="addition">Addition</option>
        <option value="subtraction">Subtraction</option>
        <option value="multiplication">Multiplication</option>
        <option value="division">Division</option>
        <option value="square_root">Square root</option>
        <option value="random_string">Random string</option>
      </select>
      <div
        class="grid grid-cols-2 gap-4"
        v-if="
          ['addition', 'subtraction', 'multiplication', 'division'].includes(
            operationType
          )
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
        v-if="['square_root'].includes(operationType)"
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
        <button class="btn btn-ghost" @click="emit('close', $event)">
          Cancel
        </button>
        <button class="btn btn-primary" @click="createOperation">Save</button>
      </div>
    </div>
  </div>
</template>
