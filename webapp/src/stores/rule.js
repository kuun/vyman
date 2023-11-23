import {ref} from 'vue'
import {defineStore} from 'pinia'

export const useRuleStore = defineStore('counter', () => {
    const selectedRuleSet = ref({})

    function setSelectedRuleSet(group) {
        selectedRuleSet.value = group
    }

    return { selectedRuleSet, setSelectedRuleSet }
})
