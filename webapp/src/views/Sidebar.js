import Backbone from 'backbone';
import NavMenuModel from '../models/NavMenuModel';

class Sidebar extends Backbone.View {
    initialize() {
        this.el = "#sidebar"
        this.model = NavMenuModel;
        this.sidebar = $().w2sidebar({
            name: 'sidebar',
            nodes: this.model.getMenus(),
            onClick: (event) => this.onClickMenu(event)
        });
    }

    onClickMenu(event) {
        if (event.node.nodes && event.node.nodes.length > 0) {
            // this node has child node, try to toggle the mode.
            this.sidebar.toggle(event.target);
        } else {
            this.model.setActiveMenuId(event.target);
        }
        // console.log(event);
    }
}

export default Sidebar;