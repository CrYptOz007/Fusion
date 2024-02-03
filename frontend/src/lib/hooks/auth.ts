import axios from 'axios';

const useRegister = async (formData: URLSearchParams) => {
	const response = await axios.post(
		'/api/user/register',
		{
			username: formData.get('username'),
			password: formData.get('password')
		},
		{
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded'
			}
		}
	);
	return response.data;
};

const useLogin = async (formData: URLSearchParams) => {
	const response = await axios.post(
		'/api/auth/login',
		{
			username: formData.get('username'),
			password: formData.get('password')
		},
		{
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded'
			},
			withCredentials: true
		}
	);
	return response.data;
};

export { useRegister, useLogin };
