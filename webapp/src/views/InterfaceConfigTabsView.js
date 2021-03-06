import Backbone from 'backbone';

class InterfaceConfigTabsView extends Backbone.View {
    preinitialize({parentLayout}) {
        this.parentLayout = parentLayout;
        this.name = 'interfaceConfigLayout';
    }

    render() {
        this.layout = $().w2layout( {
            name: this.name,
 
            panels: [
                {
                    type: 'main', 
                    tabs: [{
                        id: 'interfaceDetail',
                        caption: '网卡信息'
                    }, {
                        id: 'ipConfig',
                        caption: 'IP配置'
                    }]
                }
            ]
        });

        w2ui[this.parentLayout].content('main', this.layout);
    }

    remove() {
        this.layout.destroy();
    }
}

export default InterfaceConfigTabsView;