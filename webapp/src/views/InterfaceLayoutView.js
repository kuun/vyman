import Backbone from 'backbone';
import Interfaces from '../models/Interfaces';
import InterfaceConfigTabsView from './InterfaceConfigTabsView';
import InterfaceNavView from './InterfaceNavView';

class InterfaceLayoutView extends Backbone.View {
    preinitialize({parentLayout}) {
        this.parentLayout = parentLayout;
        this.name = 'interfacesLayout';
        this.interfacesModel = Interfaces;
    }

    render() {
        this.layout = $().w2layout({
            name: this.name,
            panels: [
                {
                    type: 'left', size: 100, resizable: true
                }, {
                    type: 'main'
                }
            ],
        });
        w2ui[this.parentLayout].content('main', this.layout);
        // create interface nav menu.
        this.interfaceNav = new InterfaceNavView({parentLayout: this.name});
        this.interfaceNav.render();
        this.interfaceNav.fetchInterfaces();
        
        this.interfaceConfigTabs = new InterfaceConfigTabsView({parentLayout: this.name});
        this.interfaceConfigTabs.render()
    }

    remove() {
        this.interfaceNav.remove();
        this.interfaceConfigTabs.remove();
        this.layout.destroy();
    }
}

export default InterfaceLayoutView;