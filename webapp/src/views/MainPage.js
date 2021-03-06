import Backbone from 'backbone';
import Sidebar from './Sidebar';
import NavMenuModel from '../models/NavMenuModel';
import InterfaceLayoutView from './InterfaceLayoutView';

class MainPage extends Backbone.View {
    initialize() {
        this.name = 'mainPageLayout';
        this.menuModel = NavMenuModel;
        this.sidebar = new Sidebar();
        this.currentView = null;
        this.listenTo(this.menuModel, 'change:activeMenuId', () => this.onMenuChanged());
    }

    render() {
        var pstyle = 'border: 1px solid #dfdfdf; padding: 5px;';
        this.layout = $('#mainPage').w2layout({
            name: this.name,
            panels: [
                {
                    type: 'top', size: 50, style: pstyle, content: 'top'
                }, {
                    type: 'left', size: 200, resizable: true
                }, {
                    type: 'main', style: pstyle, content: 'main'
                }, {
                    type: 'bottom', size: 24, resizable: false, content: 'right'
                }
            ],
        });
        this.layout.content('left', this.sidebar.sidebar);
    }

    onMenuChanged() {
        if (this.currentView) this.currentView.remove();

        var menuId = this.menuModel.getActiveMenuId();
        var view;
        
        switch (menuId) {
            case 'interfaces-ethernet':
                view = new InterfaceLayoutView({parentLayout: this.name});
        }
        view.render();
        this.currentView = view;
    }
}

export default MainPage;