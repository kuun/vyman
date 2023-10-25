import {defineStore} from 'pinia'

export const useLoginStore = defineStore('login', {
    state: () => ({
        status: 'unauthenticated',
    }),

    getters: {
        isAuthenticated() {
            return this.status === 'authenticated'
        }
    },

    actions: {
        login() {
            this.status = 'authenticated'
        },
        logout() {
            this.status = 'unauthenticated'
        },
        timeout() {
            this.status = 'unauthenticated'
        }
    }
})
