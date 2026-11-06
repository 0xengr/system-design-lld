interface Car {
  startEngine(): void;
  shiftGear(gear: number): void;
  accelerate(): void;
  brake(): void;
  stopEngine(): void;
}

/*
Concrete Class: Provides the "HOW".
Callers hold a `Car` reference and never depend on SportsCar directly
*/
class SportsCar implements Car {
  private isEngineOn = false;
  private currentSpeed = 0;
  private currentGear = 0;

  constructor(
    private readonly brand: string,
    private readonly model: string,
  ) {}

  private log(message: string): void {
    console.log(`${this.brand} ${this.model} : ${message}`);
  }

  startEngine(): void {
    this.isEngineOn = true;
    this.log("Engine starts with a roar!");
  }

  shiftGear(gear: number): void {
    if (!this.isEngineOn) {
      this.log("Engine if off! Cannot shift gear.");
      return;
    }
    this.currentGear = gear;
    this.log(`Shifted to Gear: ${this.currentGear}`);
  }

  accelerate(): void {
    if (!this.isEngineOn) {
      this.log("Engine is Off! Cannot accelerate");
      return;
    }

    this.currentSpeed += 20;
    this.log(`Accelerating to ${this.currentSpeed} km/h`);
  }

  brake(): void {
    this.currentSpeed = Math.max(0, this.currentSpeed - 20);
    this.log(`Braking! Speed is now ${this.currentSpeed} km/h`);
  }

  stopEngine(): void {
    this.isEngineOn = false;
    this.currentGear = 0;
    this.currentSpeed = 0;
    this.log("Engine turned Off.");
  }
}

// Main
const myCar: Car = new SportsCar("Ford", "Mustang");

myCar.startEngine();
myCar.shiftGear(1);
myCar.accelerate();
myCar.shiftGear(2);
myCar.accelerate();
myCar.brake();
myCar.stopEngine();

// OUTPUT:
// bun run 02-OOPs/abstraction.ts
// Ford Mustang : Engine starts with a roar!
// Ford Mustang : Shifted to Gear: 1
// Ford Mustang : Accelerating to 20 km/h
// Ford Mustang : Shifted to Gear: 2
// Ford Mustang : Accelerating to 40 km/h
// Ford Mustang : Braking! Speed is now 20 km/h
// Ford Mustang : Engine turned Off.
