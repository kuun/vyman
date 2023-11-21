import {ref} from 'vue'
import {defineStore} from 'pinia'

export const useGroupStore = defineStore('counter', () => {
  const selectedGroup = ref({})

    function setSelectedGroup(group) {
      selectedGroup.value = group
    }

  return { selectedGroup, setSelectedGroup }
})
