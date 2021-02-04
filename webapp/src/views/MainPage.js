import Backbone from 'backbone';
import Sidebar from './Sidebar';
import NavMenuModel from '../models/NavMenuModel';

class MainPage extends Backbone.View {
    initialize() {
        this.el = '#mainPage';
        this.menuModel = NavMenuModel;
        this.sidebar = new Sidebar();
        this.listenTo(this.menuModel, 'change:activeMenuId', this.onMenuChanged);
    }

    render() {
        var pstyle = 'border: 1px solid #dfdfdf; padding: 5px;';
        this.layout = $('#mainPage').w2layout({
            name: 'mainPage',
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
        this.layout.content('main', this.menuModel.getActiveMenuId());
    }
}

export default MainPage;