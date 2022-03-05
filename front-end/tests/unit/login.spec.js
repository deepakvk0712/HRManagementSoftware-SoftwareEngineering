import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import LoginPage from '../../src/components/LoginPage.vue'

describe('LoginPage.vue', () => {
    it('Check if Login is displayed ', () => {
        const wrapper = shallowMount(LoginPage);
        expect(wrapper.find('#regLoginDisplay').text()).equal('Login');
        // expect(wrapper.find('#forgot-password').text()).equal('Change Password');
        
    });

    it('Check the input fields ', () => {
        const wrapper = shallowMount(LoginPage);
        expect(wrapper.find('#userNameInput').exists()).to.exist;
        expect(wrapper.find('#userPasswordInput').exists()).to.exist;
        
    });

    it('Checking the Submit button and the text on button', () => {
        const wrapper = shallowMount(LoginPage);
        expect(wrapper.find('#regSubmitButton').exists()).to.exist;
        
        expect(wrapper.find('#regSubmitButton').text()).equal('Login');
    });
})

// // import { mount } from '@vue/test-utils'
// // import App from '/HRManagementSoftware-SoftwareEngineering/front-end/src/App.vue'


// // describe('Mounted App', () => {
// //     const wrapper = mount(App);
  
// //     test('does a wrapper exist', () => {
// //       expect(wrapper.exists()).toBe(true)
// //     })
// //   })

// // import { mount } from '@vue/test-utils'
// // import LoginPage from './LoginPage.vue'

// // describe('LoginPage.vue', ()  => {
// //   test('setup correctly', () =>{
// //     const wrapper = shallowMount(Login)
// //     expect(true).toBe(true)
// //   })
// // })

// // import { expect } from 'chai'
// // import { shallowMount } from '@vue/test-utils'
// // import LoginPage from '../../src/LoginPage.vue'

// // describe('LoginPage.vue', () => {
// //     it('If the Submit buttom works and clicks ', () => {
// //         const wrapper = shallowMount(LoginPage);
// //         expect(wrapper.find('#finCardHead').text()).equal('Banking Information');
// //     });

// //     it('Renders the 3 input fields - bank name, routing number and account number ', () => {
// //         const wrapper = shallowMount(FinancialFormPopup);
// //         expect(wrapper.find('#bankNameInput').exists()).to.exist;
// //         expect(wrapper.find('#routingNumberInput').exists()).to.exist;
// //         expect(wrapper.find('#accountNumberInput').exists()).to.exist;
// //     });

// //     it('Renders the Cancel and Submit buttons and button text ', () => {
// //         const wrapper = shallowMount(FinancialFormPopup);
// //         expect(wrapper.find('#finCancelButton').exists()).to.exist;
// //         expect(wrapper.find('#finSubmitButton').exists()).to.exist;
        
// //         expect(wrapper.find('#finCancelButton').text()).equal('Cancel');
// //         expect(wrapper.find('#finSubmitButton').text()).equal('Submit');
// //     });
// // })

// import {mount} from "@vue/test-utils";
// import LoginPage from "@/components/LoginPage.vue"

// describe("LoginPage.vue", () =>{

// it("renders component", () => {
//     const wrapper = mount (Counter);
//     expect(wrapper.element).toMatchSnapshot();

//     });

// });

