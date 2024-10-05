class DencoSeca extends Developer {
  #scaredOf;
  #headshots;

  constructor() {
    super();
    this.firstName = 'Leon';
    this.lastName = 'Brown';
    this.#headshots = {
      naughtiesBoyband: '🧑‍🦱',
      tastefulCableknit: '💁‍♂️',
    };
    this.portfolio = 'https://www.leonbrown.dev';
    this.home = 'Edinburgh, UK';
    this.#scaredOf =
      'Spiders that dissapear when you look away for like, ONE second.';
    this.isAGILE = true;
    this.skills = [
      'HTML',
      'CSS',
      'JavaScript',
      'React',
      'Redux',
      'SASS',
      'Express',
      'Node',
      'Git',
      'Docker',
      'Java',
      'DevOps',
      'Kubernetes',
      'CI/CD',
      'PSM1 Scrum Master',
      'AWS Cloud Practitioner',
    ];
  }

  isCompatibleWithJob({ description, requiredSkills, companyAddress }) {
    if (description.includes(this.#scaredOf)) {
      console.warning(`ALERT! ${companyAddress} is now a NO-GO zone.`);
      return false;
    }

    const matchedSkills = requiredSkills.map(item => this.skills.includes(item));
    matchedSkills.length >= 4
      ? console.log('Bills paid 💷')
      : console.log(
        'Apply anyway and learn fast! 👍',
      );

    return true;
  }

  applyForJob({ jobTitle, companyName, companyUrl }) {
    const useTooMuchHairGel = dictionary.some(
      word => word.replace('s', 'z') === companyName,
    );

    const application = {
      name: this.firstName + ' ' + this.lastName,
      skills: this.skills,
      photo: useTooMuchHairGel
        ? this.#headshots.naughtiesBoyband
        : this.#headshots.tastefulCableknit,
    };

    super.shamelesslyPromote({
      method: 'post',
      url: `${companyUrl}/careers/apply?position=${jobTitle}`,
      data: application,
    });
  }
}

const yourNewBestFriend = new DencoSeca();

yourNewBestFriend.isCurrentlyLearning = ['AWS Solutions Architect', 'Go'];
