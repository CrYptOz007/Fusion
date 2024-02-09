import api from '$lib/middlewares/api';

const getInfo = async (id: number) => {
	const response = await api.get(`/services/ipmi/info?id=${id}`);
	return response;
};

export default { getInfo };
