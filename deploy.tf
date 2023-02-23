terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.53"
    }
  }
}

provider "google" {
  project = "tidy-federation-377908"
  region  = "europe-west3"
}

resource "google_compute_network" "network" {
  name                    = "liqo-network"
  auto_create_subnetworks = true
}

resource "google_compute_firewall" "firewall" {
  name    = "liqo-firewall"
  network = google_compute_network.network.name

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  source_ranges = ["0.0.0.0/0"]
}

resource "google_compute_instance" "instance" {
  name         = "liqo-instance"
  machine_type = "e2-micro"
  zone         = "europe-west1-c"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = google_compute_network.network.self_link
  }

  metadata_startup_script = <<-SCRIPT
      # Install dependencies
      apt-get update
      apt-get install -y curl git

      # Install Go
      curl -sSL https://dl.google.com/go/go1.20.1.linux-amd64.tar.gz | tar -C /usr/local -xz
      echo 'export PATH=$PATH:/usr/local/go/bin' >> /root/.bashrc
    SCRIPT
}
