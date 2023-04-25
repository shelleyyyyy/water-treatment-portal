import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'Docs',
    Svg: require('@site/static/img/undraw_docusaurus_mountain.svg').default,
    description: (
      <>
        This documentation provides information about the Water Treatment Test Bed project, which is designed to simulate and test a water treatment facility using a combination of physical hardware, virtual reality, software management, and automation scripts.
      </>
    ),
  },
  {
    title: 'About',
    Svg: require('@site/static/img/undraw_docusaurus_tree.svg').default,
    description: (
      <>
       Our project includes a digital twin, which simulates IoT devices used in the water treatment facility and is built using Mininet and Neo4j. We use Docker and Docker Compose to run multiple instances of the digital twin simultaneously and Portainer to manage the containers.
      </>
    ),
  },
  {
    title: 'What this is',
    Svg: require('@site/static/img/undraw_docusaurus_react.svg').default,
    description: (
      <>
        This documentation will guide you through the setup and configuration of the various components in our test bed and provide information on how to use them. We hope this project will serve as a valuable resource for those interested in water treatment facility testing and IoT device simulation.
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
