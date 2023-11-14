import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useIfaceStore = defineStore('iface', () => {
    const selectedIface = ref({});

    const getSelectedIface = computed(() => {
        return selectedIface.value
    });

    function setSelectedIface(iface) {
        selectedIface.value = iface
    }


    return { selectedIface, getSelectedIface, setSelectedIface };
})
