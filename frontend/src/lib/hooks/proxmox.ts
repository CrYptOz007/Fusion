import api from '$lib/middlewares/api';

const getNodes = async (id: number) => {
	const response = await api.get(`/services/proxmox/nodes?id=${id}`);
	return response;
};

export default { getNodes };
