import api from '$lib/middlewares/api';

const getServices = async () => {
	const response = await api.get('/services');
	return response.data;
};

const addService = async (formData: FormData) => {
	// console log all formData fields
	const response = await api.post('/services/create', formData, {
		headers: {
			'Content-Type': 'multipart/form-data'
		}
	});
	return response;
};

const pingService = async (id: number) => {
	const response = await api.get(`/services/ping?id=${id}`);
	return response;
};

export { getServices, addService, pingService };
