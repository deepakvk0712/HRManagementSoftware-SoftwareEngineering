import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import ChangePassword from '../../src/components/ChangePassword.vue'

describe('ChangePassword.vue', () => {
    // it('Check if Change Password heading is displayed ', () => {
    //     const wrapper = shallowMount(ChangePassword);
    //     expect(wrapper.find('#regChangePassword').text()).equal("Change your Password here!");
    //     // expect(wrapper.find('#forgot-password').text()).equal('Change Password');
        
    // });

    it('Check the input fields ', () => {
        const wrapper = shallowMount(ChangePassword);
        expect(wrapper.find('#regEmailID').exists()).to.exist;
        expect(wrapper.find('#regOldPassword').exists()).to.exist;
        expect(wrapper.find('#regNewPassword').exists()).to.exist;

        
    });

    // it('Checking the Submit and Clear button and the text on button', () => {
    //     const wrapper = shallowMount(ChangePassword);
    //     expect(wrapper.find('#regSubmitButton').exists()).to.exist;
    //     expect(wrapper.find('#regClearButton').exists()).to.exist;

        
    //     expect(wrapper.find('#regSubmitButton').text()).equal('Submit');
    //     expect(wrapper.find('#regClearButton').text()).equal('Clear');
    // });

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
// // import ChangePassword from './ChangePassword.vue'

// // describe('ChangePassword.vue', ()  => {
// //   test('setup correctly', () =>{
// //     const wrapper = shallowMount(Login)
// //     expect(true).toBe(true)
// //   })
// // })

// // import { expect } from 'chai'
// // import { shallowMount } from '@vue/test-utils'
// // import ChangePassword from '../../src/ChangePassword.vue'

// // describe('ChangePassword.vue', () => {
// //     it('If the Submit buttom works and clicks ', () => {
// //         const wrapper = shallowMount(ChangePassword);
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
// import ChangePassword from "@/components/ChangePassword.vue"

// describe("ChangePassword.vue", () =>{

// it("renders component", () => {
//     const wrapper = mount (Counter);
//     expect(wrapper.element).toMatchSnapshot();

//     });

// });

