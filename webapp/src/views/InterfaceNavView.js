import Backbone from 'backbone';
import Interfaces from '../models/Interfaces';
import _ from 'lodash';

class InterfaceNavView extends Backbone.View {
    preinitialize({parentLayout}) {
        this.name = 'interfacesNav';
        this.interfacesModel = Interfaces;
        this.parentLayout = parentLayout;
    }

    initialize() {
        this.interfacesModel.listenTo(this.interfacesModel, "change:interfaces", ()=>this.onInterfacesChanged()); 
   
    }

    render() {
        var parentLayout = w2ui[this.parentLayout];
        this.navBar = $().w2sidebar({
            name: this.name,
            topHTML: '<div>网卡列表</div>',
            nodes: this.buildSidebarNodes(this.interfacesModel.getInterfaces())
        });
        parentLayout.content('left', this.navBar);
    }

    remove() {
        this.navBar.distroy();
    }

    onInterfacesChanged() {
        // remove old interfaces in list.
        _.each(this.navBar.nodes, (node) => {
            this.navBar.remove(node.id);
        });

        var ifaces = this.interfacesModel.getInterfaces();
        var nodes = this.buildSidebarNodes(ifaces) ;
        this.navBar.add(nodes);
        this.navBar.select(nodes[0].id);
    }

    buildSidebarNodes(interfaces) {
        return _.map(interfaces, (iface) => {
            return {
                id: iface.name,
                text: iface.name,
                icon: 'fa fa-star-o'
            };
        });
    }

    fetchInterfaces() {
        $.getJSON('/api/interfaces/ethernet', (data) => {
            this.interfacesModel.setInterfaces(data);
        }).fail((res) => {
            alert('获取网卡列表失败!');
        })
    }
}

export default InterfaceNavView;