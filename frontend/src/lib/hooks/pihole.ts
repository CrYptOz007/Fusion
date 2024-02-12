import api from '$lib/middlewares/api';

const getSummary = async (id: number) => {
	const response = await api.get(`/services/pihole/summary?id=${id}`);
	return response;
};

const getSummaryData = async (id: number) => {
	const response = await api.get(`/services/pihole/summary?id=${id}`);
	return response.data;
};

const enable = async (id: number) => {
	const response = await api.get(`/services/pihole/enable?id=${id}`);
	return response;
};

const disable = async (id: number) => {
	const response = await api.get(`/services/pihole/disable?id=${id}`);
	return response;
};

export default { getSummary, getSummaryData, enable, disable };
