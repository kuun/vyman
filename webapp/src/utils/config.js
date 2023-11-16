import axios from "axios";

export default {
    command(op, path, value) {
        return {
            op,
            path: path.split(' '),
            value
        };
    },

    configure(cmds) {
        return axios.post("/api/configure", cmds)
    },

    showConfig(path) {
        return axios.post("/api/showConfig", {path})
    },

    returnValue(path) {
        return axios.post("/api/returnValue", {path})
    },

    returnValues(path) {
        return axios.post("/api/returnValues", {path})
    }
}