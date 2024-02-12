import api from '$lib/middlewares/api';

const getNodes = async (id: number) => {
	const response = await api.get(`/services/proxmox/nodes?id=${id}`);
	return response;
};

const getVMs = async (id: number) => {
	const nodes: { node: string; status: string }[] = (await getNodes(id)).data.data;

	const data = await Promise.all(
		nodes.map(async (node) => {
			if (node.status !== 'online') return;
			const response = await api.get(
				`/services/proxmox/virtualmachines?id=${id}&node=${node.node}`
			);
			return response.data.data;
		})
	);

	return data.flat();
};

const getContainers = async (id: number) => {
	const nodes: { node: string; status: string }[] = (await getNodes(id)).data.data;

	const data = await Promise.all(
		nodes.map(async (node) => {
			if (node.status !== 'online') return;
			const response = await api.get(`/services/proxmox/containers?id=${id}&node=${node.node}`);
			return response.data.data;
		})
	);

	return data.flat();
};

export default { getNodes, getVMs, getContainers };
