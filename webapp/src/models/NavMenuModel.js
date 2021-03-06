import Backbone from 'backbone';

class NavMenuModel extends Backbone.Model {
    defaults() {
        return {
            menus: [{
                id: 'interfaces',
                text: '网口管理',
                icon: 'fa fa-internet-explorer',
                nodes: [{
                    id: 'interfaces-ethernet',
                    text: '以太网',
                    icon: 'fa fa-cog'
                }]
            }, {
                id: 'groups',
                text: '组管理',
                icon: 'fa fa-sitemap',
                nodes: [{
                    id: 'address-group',
                    text: '地址组',
                    icon: 'fa fa-cog'
                }, {
                    id: 'network-group',
                    text: '网络组',
                    icon: 'fa fa-cog'
                }, {
                    id: 'port-group',
                    text: '端口组',
                    icon: 'fa fa-cog'
                }]
            }, {
                id: 'rule-sets',
                text: '规则集管理',
                icon: 'fa fa-list-ul'
            }, {
                id: 'zone',
                text: '安全域',
                icon: 'fa fa-th-large'
            }, {
                id: 'log',
                text: '审计日志',
                icon: 'fa fa-bar-chart'
            }],
            activeMenuId: null,
        };
    }

    getMenus() {
        return this.get('menus');
    }

    getActiveMenuId() {
        return this.get('activeMenuId')
    }

    setActiveMenuId(activeMenuId) {
        this.set({ activeMenuId })
    }
}

export default new NavMenuModel()