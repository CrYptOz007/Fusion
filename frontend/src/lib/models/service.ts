import ipmi from '$lib/hooks/ipmi';
import pihole from '$lib/hooks/pihole';
import proxmox from '$lib/hooks/proxmox';
import { pingService } from '$lib/hooks/services';

export class Service {
	public id: number;
	public type: string;
	public name: string;
	public icon: string;
	public hostname: string;
	public port: number;
	public serviceType: ServiceType;
	public online: boolean;

	constructor(service: {
		ID: number;
		type: string;
		name: string;
		icon: string;
		hostname: string;
		port: number;
	}) {
		this.id = service.ID;
		this.type = service.type;
		this.name = service.name;
		this.icon = service.icon;
		this.hostname = service.hostname;
		this.port = service.port;
		this.serviceType = this.setType();
		this.online = false;
	}

	private setType(): ServiceType {
		switch (this.type) {
			case 'proxmox':
				return new ProxmoxService(this.id);
			case 'pihole':
				return new PiholeService(this.id);
			case 'ipmi':
				return new IpmiService(this.id);
			default:
				return new GenericService(this.id);
		}
	}

	public goToService(): void {
		if (this.port === 80) window.open(`http://${this.hostname}`);
		else {
			window.open(`https://${this.hostname}:${this.port}`);
		}
	}
}

abstract class ServiceType {
	protected id: number;

	constructor(id: number) {
		this.id = id;
	}

	public abstract executePing(): Promise<boolean>;
}

export class GenericService extends ServiceType {
	public async executePing(): Promise<boolean> {
		return (await pingService(this.id)).status === 200;
	}
}

export class ProxmoxService extends ServiceType {
	public async executePing(): Promise<boolean> {
		return (await proxmox.getNodes(this.id)).status === 200;
	}
}

export class PiholeService extends ServiceType {
	public async executePing(): Promise<boolean> {
		return (await pihole.getSummary(this.id)).status === 200;
	}
}

export class IpmiService extends ServiceType {
	public async executePing(): Promise<boolean> {
		return (await ipmi.getInfo(this.id)).status === 200;
	}
}
