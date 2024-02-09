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

const useRefresh = async () => {
	const response = await axios.get('/api/auth/refresh');
	return response;
};

const useLogout = async () => {
	const response = await axios.get('/api/auth/logout');
	return response;
};

export { useRegister, useLogin, useRefresh, useLogout };
