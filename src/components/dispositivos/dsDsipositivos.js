export default class dsDispositivo {
  constructor() {
    this.dispositivos = [
      {
      }
    ];
  }

  add(dispositivo) {
    return new Promise((resolve) => {
      this.dispositivos.push(dispositivo);
      resolve(dispositivo);
    });
  }

  getAll() {
    return new Promise((resolve) => {
      resolve(this.dispositivos);
    });
  }
}
