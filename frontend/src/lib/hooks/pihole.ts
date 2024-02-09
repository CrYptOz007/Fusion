import api from '$lib/middlewares/api';

const getSummary = async (id: number) => {
	const response = await api.get(`/services/pihole/summary?id=${id}`);
	return response;
};

export default { getSummary };
