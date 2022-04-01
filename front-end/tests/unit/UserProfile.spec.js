import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import UserProfileView from '../../src/views/UserProfileView.vue'

describe('UserProfileView.vue', () => {
    it('Renders the User Profile Heading ', () => {
        const wrapper = shallowMount(UserProfileView);
        expect(wrapper.find('#usrProfHead').text()).equal('Account Details');
    });

    it('Renders Update form fields correctly ', () => {
        const wrapper = shallowMount(UserProfileView);
        expect(wrapper.find('#usrProFName').text()).equal('First Name');
        expect(wrapper.find('#usrProLName').text()).equal('Last Name');
        expect(wrapper.find('#usrPhoneNumber').text()).equal('Phone Number');
        expect(wrapper.find('#usrEmail').text()).equal('Email Address');
        expect(wrapper.find('#usrAddr').text()).equal('Address');
        expect(wrapper.find('#usrCity').text()).equal('City');
        expect(wrapper.find('#usrState').text()).equal('State');
        expect(wrapper.find('#usrZip').text()).equal('Zip');
        expect(wrapper.find('#usrDesc').text()).equal('Description');
    });

    it('Renders the Update button in the user profile page ', () => {
        const wrapper = shallowMount(UserProfileView);
        expect(wrapper.find('#usrUpdateButton').exists()).to.exist;
        
        expect(wrapper.find('#usrUpdateButton').text()).equal('Update');
    });

    it('Renders the different fields for updating the user information ', () => {
        const wrapper = shallowMount(UserProfileView);
        expect(wrapper.find('#usrProDescInput').exists()).to.exist;
        expect(wrapper.find('#usrProZipInput').exists()).to.exist;
        expect(wrapper.find('#usrProStateInput').exists()).to.exist;
        expect(wrapper.find('#usrProCityInput').exists()).to.exist;
        expect(wrapper.find('#usrProAddressInput').exists()).to.exist;
        expect(wrapper.find('#usrProEmailInput').exists()).to.exist;
        expect(wrapper.find('#usrProNumberInput').exists()).to.exist;
        expect(wrapper.find('#usrProLNameInput').exists()).to.exist;
        expect(wrapper.find('#usrProFNameInput').exists()).to.exist; 
    });
    
    it('Correctly renders the productivity header and productivity score ', () => {
        const wrapper = shallowMount(UserProfileView);
        expect(wrapper.find('#prodUserProfile').exists()).to.exist;
    });

    it('Correctly renders the about me and the description sent from backend ', () => {
        const wrapper = shallowMount(UserProfileView);
        expect(wrapper.find('#aboutUserProfile').exists()).to.exist;
    });

    it('Correctly renders the team members header and the team members email ID and name ', () => {
        const wrapper = shallowMount(UserProfileView);
        expect(wrapper.find('#teamMembersUserProfile').exists()).to.exist;
    });
})