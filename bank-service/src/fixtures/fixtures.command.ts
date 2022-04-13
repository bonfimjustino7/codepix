import { Command, Console } from 'nestjs-console';
import { getConnection } from 'typeorm';
import fixtures from './fixtures';
import * as chalk from 'chalk';

@Console()
export class FixturesCommand {
  @Command({
    command: 'fixtures',
    description: 'Seed data in database',
  })
  async command() {
    await this.runMigrations();
    for (const fixture of fixtures) {
      await this.createInDatabase(fixture.model, fixture.fields);
    }

    console.log(chalk.green('Data generated'));
    process.exit();
  }

  async runMigrations() {
    const conn = getConnection('default');

    for (const migration of conn.migrations.reverse()) {
      await conn.undoLastMigration();
    }
  }

  async createInDatabase(model: any, data: any) {
    const respository = this.getRepository(model);
    const obj = respository.create(data);
    await respository.save(obj);
  }

  getRepository(model: any) {
    const conn = getConnection('default');
    return conn.getRepository(model);
  }
}
