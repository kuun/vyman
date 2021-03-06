import Backbone from 'backbone';

class Interfaces extends Backbone.Model {
    defaults() {
        return {
            'interfaces': [],
            'selectedIface': ''
        };
    }

    setInterfaces(interfaces) {
        this.set({interfaces});
    }

    getInterfaces() {
        return this.get('interfaces');
    }

    setSelectedIface(selectedIface) {
        this.set({selectedIface});
    }

    getSelectedIface() {
        return this.get('selectedIface');
    }
}

export default new Interfaces();