<template>
    <nav :style="{background: background || '#444'}">
        <ul :style="{background: background || '#444'}" ref="nav">
            <figure class="logo" @click="toggleNav">
                <img :src = "imagePath" width="40" height="40" padding-right="10"/>
            </figure>
            <li
                v-for="(tab, index) in tabInfo"
                :key="index"
                @mouseenter = "$event.currentTarget.style.background = hoverColor || '#999'"
                @mouseleave = "$event.currentTarget.style.background = background || '#444'"
            >
                <router-link 
                    :to="tab.path"
                    :style="{color: tabColor || '#BBB'}"
                >
                    {{tab.text}}
                    <i style="padding-right: 8px;" :class = "tab.icon"/>
                </router-link>
            </li>
        </ul>
    </nav>
</template>

<script>

export default {
    props : ['tabInfo', 'background', 'tabColor', 'imagePath'],
    methods : {
        toggleNav () {
            const ns = this.$refs.nav.classList
            ns.contains('active') ? ns.remove('active') : ns.add('active')
        }
    }
}

</script>

<style scoped lang="scss">
nav{
    height : 70px;
    width : 100%;
    ul{
        display: flex;
        height: 100%;
        align-items: center;
        margin-block-start: 0;
        margin-block-end: 0;
        padding-inline-start: 0;

        a{
            text-decoration: none;
            display: flex;
            flex-direction: row-reverse;
            align-items: center;
        }

        i{
            font-size: 26px
        }

        li{
            list-style-type: none;
            padding: 10px 20px;
        }
        figure{
            cursor: pointer;
        }
    }
}

@media screen and (max-width: 759px) {
    nav{
        ul{
            position: absolute;
            width: 300px;
            flex-direction: column;
            left: -230px;
            top: 60px;
            transition: 300ms ease all;

            &.active{
                left: 0;
            }

            figure{
                position: fixed;
                z-index: 1;
                top: 10px;
                left: 5px;
            }

            li{
                width: 100%;
                padding-left: 0px;
                padding-right: 0px;
            }

            a{
                flex-direction: row;
                margin-left: 20px;
                justify-content: space-between;
                margin-right: 20px;
            }
        }
    }
}
</style>>